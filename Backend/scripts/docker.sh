#!/bin/bash

# Interactive Docker compose wrapper

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Available environments
ENVS=("local" "dev" "prod")

# Available actions
ACTIONS=("up" "down" "build" "logs" "restart" "ps")

echo -e "${BLUE}=== Docker Environment Manager ===${NC}"
echo ""

# Select environment
echo "Select environment:"
for i in "${!ENVS[@]}"; do
    echo "  $((i+1)). ${ENVS[$i]}"
done
read -p "Enter choice [1-${#ENVS[@]}]: " env_choice

if [[ ! "$env_choice" =~ ^[1-${#ENVS[@]}]$ ]]; then
    echo -e "${RED}Invalid choice${NC}"
    exit 1
fi

ENV="${ENVS[$((env_choice-1))]}"
COMPOSE_FILE="docker/docker-compose.${ENV}.yml"

if [ ! -f "$COMPOSE_FILE" ]; then
    echo -e "${RED}Compose file not found: $COMPOSE_FILE${NC}"
    exit 1
fi

echo ""
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
echo -e "${YELLOW}Running: docker compose -f $COMPOSE_FILE $ACTION${NC}"
echo ""

case $ACTION in
    "up")
        docker compose -f "$COMPOSE_FILE" up -d
        ;;
    "down")
        docker compose -f "$COMPOSE_FILE" down
        ;;
    "build")
        docker compose -f "$COMPOSE_FILE" build
        ;;
    "logs")
        docker compose -f "$COMPOSE_FILE" logs -f
        ;;
    "restart")
        docker compose -f "$COMPOSE_FILE" restart
        ;;
    "ps")
        docker compose -f "$COMPOSE_FILE" ps
        ;;
esac

echo ""
echo -e "${GREEN}Done!${NC}"
