#!/bin/bash

# Interactive migration runner

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Load env if exists
if [ -f .env ]; then
    export $(cat .env | grep -v '#' | xargs)
fi

DB_HOST=${DB_HOST:-127.0.0.1}
DB_PORT=${DB_PORT:-3306}
DB_USER=${DB_USER:-root}
DB_PASSWORD=${DB_PASSWORD:-}
DB_NAME=${DB_NAME:-be_db}
MIGRATION_DIR="./migrations/versions"

DB_URL="${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true"

echo -e "${BLUE}=== Migration Runner ===${NC}"
echo -e "Database: ${DB_NAME}@${DB_HOST}:${DB_PORT}"
echo ""

# Available actions
ACTIONS=("up" "down" "status" "version" "create" "reset")

echo "Select action:"
for i in "${!ACTIONS[@]}"; do
    echo "  $((i+1)). ${ACTIONS[$i]}"
done
read -p "Enter choice [1-${#ACTIONS[@]}]: " action_choice

if [[ ! "$action_choice" =~ ^[1-${#ACTIONS[@]}]$ ]]; then
    echo -e "${RED}Invalid choice${NC}"
    exit 1
fi

ACTION="${ACTIONS[$((action_choice-1))]}"

echo ""
echo -e "${YELLOW}Running: goose $ACTION${NC}"
echo ""

case $ACTION in
    "up")
        goose -dir "$MIGRATION_DIR" mysql "$DB_URL" up
        ;;
    "down")
        goose -dir "$MIGRATION_DIR" mysql "$DB_URL" down
        ;;
    "status")
        goose -dir "$MIGRATION_DIR" mysql "$DB_URL" status
        ;;
    "version")
        goose -dir "$MIGRATION_DIR" mysql "$DB_URL" version
        ;;
    "create")
        read -p "Enter migration name: " name
        if [ -z "$name" ]; then
            echo -e "${RED}Error: Migration name is required${NC}"
            exit 1
        fi
        goose -dir "$MIGRATION_DIR" create "$name" sql
        ;;
    "reset")
        echo -e "${RED}WARNING: This will reset all migrations!${NC}"
        read -p "Are you sure? (yes/no): " confirm
        if [ "$confirm" = "yes" ]; then
            goose -dir "$MIGRATION_DIR" mysql "$DB_URL" reset
            goose -dir "$MIGRATION_DIR" mysql "$DB_URL" up
        else
            echo "Cancelled."
        fi
        ;;
esac

echo ""
echo -e "${GREEN}Done!${NC}"
