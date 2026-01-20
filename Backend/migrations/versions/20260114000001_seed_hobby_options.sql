-- +goose Up

-- ============================================================================
-- SEED DATA: HOBBY OPTIONS
-- A curated list of hobbies organized by category
-- ============================================================================

INSERT INTO hobby_options (id, name, category, icon_name, micro_descriptions, is_active, sort_order) VALUES
-- Creative & Arts
(UUID(), 'Photography', 'Creative', 'camera', '["Landscape", "Portrait", "Street", "Wildlife", "Amateur"]', TRUE, 1),
(UUID(), 'Painting', 'Creative', 'palette', '["Watercolor", "Oil", "Acrylic", "Digital", "Abstract"]', TRUE, 2),
(UUID(), 'Drawing', 'Creative', 'pencil', '["Sketching", "Charcoal", "Digital Art", "Comics", "Portraits"]', TRUE, 3),
(UUID(), 'Writing', 'Creative', 'pen', '["Fiction", "Poetry", "Blogging", "Journaling", "Screenwriting"]', TRUE, 4),
(UUID(), 'Music', 'Creative', 'music', '["Playing Instruments", "Singing", "Composing", "DJing", "Listening"]', TRUE, 5),
(UUID(), 'Crafts', 'Creative', 'scissors', '["Knitting", "Pottery", "Origami", "Jewelry Making", "DIY Projects"]', TRUE, 6),

-- Sports & Fitness
(UUID(), 'Running', 'Sports', 'running', '["Marathon", "Trail Running", "Casual Jogging", "Sprinting", "Treadmill"]', TRUE, 10),
(UUID(), 'Gym & Fitness', 'Sports', 'dumbbell', '["Weight Training", "CrossFit", "HIIT", "Bodybuilding", "Cardio"]', TRUE, 11),
(UUID(), 'Yoga', 'Sports', 'yoga', '["Hatha", "Vinyasa", "Hot Yoga", "Meditation", "Beginner"]', TRUE, 12),
(UUID(), 'Swimming', 'Sports', 'water', '["Laps", "Ocean Swimming", "Water Polo", "Diving", "Casual"]', TRUE, 13),
(UUID(), 'Cycling', 'Sports', 'bicycle', '["Road Cycling", "Mountain Biking", "Commuting", "BMX", "Spinning"]', TRUE, 14),
(UUID(), 'Team Sports', 'Sports', 'ball', '["Football", "Basketball", "Cricket", "Volleyball", "Badminton"]', TRUE, 15),
(UUID(), 'Hiking', 'Sports', 'mountain', '["Day Hikes", "Trekking", "Mountain Climbing", "Nature Walks", "Backpacking"]', TRUE, 16),
(UUID(), 'Dancing', 'Sports', 'dance', '["Salsa", "Hip Hop", "Contemporary", "Bollywood", "Ballet"]', TRUE, 17),
(UUID(), 'Martial Arts', 'Sports', 'martial', '["Karate", "Judo", "MMA", "Kickboxing", "Taekwondo"]', TRUE, 18),

-- Food & Drinks
(UUID(), 'Cooking', 'Food', 'chef', '["Home Cooking", "Baking", "Grilling", "Healthy Meals", "Exotic Cuisines"]', TRUE, 20),
(UUID(), 'Coffee', 'Food', 'coffee', '["Brewing", "Cafe Hopping", "Latte Art", "Tasting", "Collection"]', TRUE, 21),
(UUID(), 'Wine & Spirits', 'Food', 'wine', '["Wine Tasting", "Cocktails", "Craft Beer", "Whiskey", "Mixology"]', TRUE, 22),
(UUID(), 'Food Exploration', 'Food', 'food', '["Street Food", "Fine Dining", "Food Photography", "Reviews", "New Cuisines"]', TRUE, 23),

-- Entertainment
(UUID(), 'Movies', 'Entertainment', 'film', '["Cinema", "Indie Films", "Documentaries", "Film Critique", "Netflix"]', TRUE, 30),
(UUID(), 'Gaming', 'Entertainment', 'gamepad', '["PC Gaming", "Console", "Mobile", "Board Games", "VR"]', TRUE, 31),
(UUID(), 'Reading', 'Entertainment', 'book', '["Fiction", "Non-Fiction", "Fantasy", "Self-Help", "Audiobooks"]', TRUE, 32),
(UUID(), 'Anime & Manga', 'Entertainment', 'tv', '["Watching Anime", "Reading Manga", "Cosplay", "Fan Art", "Conventions"]', TRUE, 33),
(UUID(), 'Podcasts', 'Entertainment', 'headphones', '["True Crime", "Comedy", "Educational", "Business", "Interviews"]', TRUE, 34),

-- Outdoors & Nature
(UUID(), 'Gardening', 'Outdoors', 'plant', '["Indoor Plants", "Vegetable Garden", "Flowers", "Balcony Garden", "Succulents"]', TRUE, 40),
(UUID(), 'Camping', 'Outdoors', 'tent', '["Backpacking", "Glamping", "RV Camping", "Wild Camping", "Family Trips"]', TRUE, 41),
(UUID(), 'Bird Watching', 'Outdoors', 'bird', '["Photography", "Identification", "Nature Walks", "Migration", "Local Parks"]', TRUE, 42),
(UUID(), 'Stargazing', 'Outdoors', 'star', '["Astronomy", "Telescope", "Night Sky Photography", "Meteor Showers", "Casual"]', TRUE, 43),

-- Travel & Adventure
(UUID(), 'Traveling', 'Travel', 'plane', '["Solo Travel", "Adventure", "Cultural", "Budget", "Luxury"]', TRUE, 50),
(UUID(), 'Road Trips', 'Travel', 'car', '["Weekend Getaways", "Cross Country", "Scenic Routes", "Camping Trips", "City Hopping"]', TRUE, 51),
(UUID(), 'Backpacking', 'Travel', 'backpack', '["Budget Travel", "Hostels", "Long-term", "Southeast Asia", "Europe"]', TRUE, 52),

-- Technology
(UUID(), 'Coding', 'Technology', 'code', '["Web Development", "Mobile Apps", "AI/ML", "Open Source", "Learning"]', TRUE, 60),
(UUID(), 'Tech Gadgets', 'Technology', 'smartphone', '["Reviews", "Early Adopter", "Smart Home", "Wearables", "Photography Gear"]', TRUE, 61),
(UUID(), 'Cryptocurrency', 'Technology', 'crypto', '["Trading", "DeFi", "NFTs", "Research", "Long-term Investing"]', TRUE, 62),

-- Social & Community
(UUID(), 'Volunteering', 'Social', 'heart', '["Animal Shelter", "Teaching", "Environment", "Community Service", "NGOs"]', TRUE, 70),
(UUID(), 'Language Learning', 'Social', 'language', '["Spanish", "Japanese", "French", "Hindi", "Mandarin"]', TRUE, 71),
(UUID(), 'Book Clubs', 'Social', 'users', '["Fiction", "Non-Fiction", "Online", "Local Meetups", "Genre-Specific"]', TRUE, 72),

-- Wellness & Mindfulness
(UUID(), 'Meditation', 'Wellness', 'zen', '["Guided", "Mindfulness", "Breathwork", "Morning Practice", "Sleep"]', TRUE, 80),
(UUID(), 'Self-Improvement', 'Wellness', 'growth', '["Productivity", "Habits", "Goal Setting", "Journaling", "Reading"]', TRUE, 81),
(UUID(), 'Spirituality', 'Wellness', 'spirit', '["Yoga Philosophy", "Buddhism", "Prayer", "Retreats", "Study"]', TRUE, 82);

-- +goose Down

DELETE FROM hobby_options;
