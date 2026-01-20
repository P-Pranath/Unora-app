# Waitlist Setup Instructions

## Free Services Setup

### 1. Supabase (Free Database) - 5 minutes

1. Go to https://supabase.com and sign up (free)
2. Create a new project
3. Go to **SQL Editor** and run this SQL to create the waitlist table:

```sql
-- Create waitlist table
CREATE TABLE waitlist (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Create index for faster lookups
CREATE INDEX idx_waitlist_email ON waitlist(email);
CREATE INDEX idx_waitlist_created_at ON waitlist(created_at DESC);

-- Enable Row Level Security (optional but recommended)
ALTER TABLE waitlist ENABLE ROW LEVEL SECURITY;

-- Create policy to allow inserts (but not reads for security)
CREATE POLICY "Allow public inserts" ON waitlist
  FOR INSERT
  TO anon
  WITH CHECK (true);
```

4. Go to **Settings** → **API**
5. Copy your **Project URL** and **anon/public key**
6. Add them to your `.env.local` file (see below)

### 2. Resend (Free Email Service) - 3 minutes

1. Go to https://resend.com and sign up (free - 3000 emails/month)
2. Go to **API Keys** and create a new key
3. Copy the API key
4. Add it to your `.env.local` file (see below)

### 3. Environment Variables

Create a `.env.local` file in your project root:

```env
# Supabase
NEXT_PUBLIC_SUPABASE_URL=https://your-project.supabase.co
NEXT_PUBLIC_SUPABASE_ANON_KEY=your-anon-key-here

# Resend
RESEND_API_KEY=your-resend-key-here

# Your app URL (for email links)
NEXT_PUBLIC_APP_URL=https://yourdomain.com
```

**Important:** Add `.env.local` to your `.gitignore` file to keep secrets safe!

### 4. Netlify Environment Variables

Since you're on Netlify:

1. Go to your Netlify dashboard
2. Navigate to **Site settings** → **Environment variables**
3. Add all the variables from `.env.local`:
   - `NEXT_PUBLIC_SUPABASE_URL`
   - `NEXT_PUBLIC_SUPABASE_ANON_KEY`
   - `RESEND_API_KEY`
   - `NEXT_PUBLIC_APP_URL`

### 5. Deploy

1. Commit and push your changes
2. Netlify will automatically redeploy
3. Test the waitlist form!

## Email Validation

The form now validates:
- ✅ Email must contain @ symbol
- ✅ Must have username before @
- ✅ Must have valid domain after @
- ✅ Domain format validation
- ✅ Checks against common email providers (Gmail, Yahoo, Hotmail, etc.)

## What Happens When Someone Submits

1. **Email validation** - Client and server-side checks
2. **Saved to Supabase** - Free database storage
3. **Confirmation email sent** - Via Resend (free tier)
4. **User sees success message** - "Check your email for confirmation!"

## Viewing Waitlist Entries

1. Go to your Supabase dashboard
2. Navigate to **Table Editor**
3. Select the `waitlist` table
4. See all emails that joined!

## Free Tier Limits

- **Supabase**: 500MB database, 2GB bandwidth/month
- **Resend**: 3,000 emails/month
- Both are perfect for a waitlist! 🎉

## Troubleshooting

### Emails not saving?
- Check Supabase environment variables are set correctly
- Check Supabase table was created properly
- Check browser console for errors

### Confirmation emails not sending?
- Check Resend API key is correct
- Check Resend dashboard for email logs
- Emails will be sent from `onboarding@resend.dev` (you can verify your domain later)

### Form not working on Netlify?
- Make sure all environment variables are set in Netlify dashboard
- Check Netlify function logs for errors
- Rebuild the site after adding environment variables

