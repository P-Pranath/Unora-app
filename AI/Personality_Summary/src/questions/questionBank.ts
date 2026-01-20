// ============================================================================
// Question Bank - Personality Intelligence Service
// ============================================================================
// Contains 40+ situational questions across all 7 personality dimensions.
// Each question presents a real-world scenario and multiple-choice options
// that reveal personality tendencies through natural responses.
// 
// DESIGN PRINCIPLES:
// - Questions are situational, never self-descriptive
// - Options don't have "right" or "wrong" answers
// - Each option has weighted impacts on 1-3 dimensions
// - Impact values range from -0.3 to +0.3 (small deltas)
// ============================================================================

import { Question, PersonalityDimension } from '../types';

// ============================================================================
// Question Bank
// ============================================================================

export const QUESTION_BANK: Question[] = [
    // ==========================================================================
    // EMOTIONAL REGULATION (6 questions)
    // ==========================================================================
    // Measures how one manages and processes emotional responses.
    // Higher scores = more regulated, processes emotions smoothly
    // Lower scores = more reactive, emotions are felt intensely

    {
        id: 'Q_ER_01',
        scenario: 'A plan you were excited about gets cancelled last minute.',
        dimensionTargets: ['emotional_regulation'],
        options: [
            {
                label: "I'm annoyed briefly, then move on to something else",
                impacts: { emotional_regulation: 0.25 }
            },
            {
                label: 'It throws me off for a while, but I adjust',
                impacts: { emotional_regulation: 0.0 }
            },
            {
                label: 'I feel quite disappointed and need some time to reset',
                impacts: { emotional_regulation: -0.2 }
            }
        ]
    },
    {
        id: 'Q_ER_02',
        scenario: 'Someone close to you gives you unexpected critical feedback.',
        dimensionTargets: ['emotional_regulation', 'conflict_posture'],
        options: [
            {
                label: 'I take a moment, then try to understand their perspective',
                impacts: { emotional_regulation: 0.2, conflict_posture: 0.1 }
            },
            {
                label: 'I feel defensive at first but come around later',
                impacts: { emotional_regulation: 0.0, conflict_posture: -0.1 }
            },
            {
                label: 'It stings and I need space before I can process it',
                impacts: { emotional_regulation: -0.15, conflict_posture: -0.15 }
            }
        ]
    },
    {
        id: 'Q_ER_03',
        scenario: 'You receive unexpectedly good news about something important.',
        dimensionTargets: ['emotional_regulation', 'energy_orientation'],
        options: [
            {
                label: "I'm pleased and feel a quiet sense of satisfaction",
                impacts: { emotional_regulation: 0.15, energy_orientation: -0.1 }
            },
            {
                label: "I'm excited and want to share with someone right away",
                impacts: { emotional_regulation: 0.0, energy_orientation: 0.2 }
            },
            {
                label: 'I feel a rush of excitement that lasts a while',
                impacts: { emotional_regulation: -0.1, energy_orientation: 0.1 }
            }
        ]
    },
    {
        id: 'Q_ER_04',
        scenario: "You're stuck in traffic and running late to something important.",
        dimensionTargets: ['emotional_regulation'],
        options: [
            {
                label: 'I accept it and use the time to listen to something',
                impacts: { emotional_regulation: 0.25 }
            },
            {
                label: "I'm frustrated but try to stay calm",
                impacts: { emotional_regulation: 0.05 }
            },
            {
                label: 'I feel stressed and keep checking the time',
                impacts: { emotional_regulation: -0.2 }
            }
        ]
    },
    {
        id: 'Q_ER_05',
        scenario: 'A project you worked hard on receives mixed reviews.',
        dimensionTargets: ['emotional_regulation', 'communication_style'],
        options: [
            {
                label: 'I focus on the constructive parts and note improvements',
                impacts: { emotional_regulation: 0.2, communication_style: 0.1 }
            },
            {
                label: "I feel a bit deflated but know it's part of the process",
                impacts: { emotional_regulation: 0.0, communication_style: 0.0 }
            },
            {
                label: "It's hard not to take the criticism personally",
                impacts: { emotional_regulation: -0.2, communication_style: -0.1 }
            }
        ]
    },
    {
        id: 'Q_ER_06',
        scenario: "Something you were counting on falls through at the last moment.",
        dimensionTargets: ['emotional_regulation', 'consistency_style'],
        options: [
            {
                label: 'I quickly look for alternatives and adapt',
                impacts: { emotional_regulation: 0.2, consistency_style: -0.15 }
            },
            {
                label: "I'm upset but start problem-solving after a bit",
                impacts: { emotional_regulation: 0.0, consistency_style: 0.0 }
            },
            {
                label: 'I feel rattled and need time to regroup',
                impacts: { emotional_regulation: -0.2, consistency_style: 0.1 }
            }
        ]
    },

    // ==========================================================================
    // COMMUNICATION STYLE (6 questions)
    // ==========================================================================
    // Measures preference for direct vs. indirect communication.
    // Higher scores = more direct, says what they mean explicitly
    // Lower scores = more indirect, prefers nuance and subtext

    {
        id: 'Q_CS_01',
        scenario: 'A friend asks if you like their new outfit, but you honestly don\'t.',
        dimensionTargets: ['communication_style', 'emotional_availability'],
        options: [
            {
                label: "I gently say it's not my favorite but they look great anyway",
                impacts: { communication_style: 0.15, emotional_availability: 0.1 }
            },
            {
                label: 'I focus on something positive about it instead',
                impacts: { communication_style: -0.15, emotional_availability: 0.15 }
            },
            {
                label: "I'm honest and tell them what I really think",
                impacts: { communication_style: 0.25, emotional_availability: -0.05 }
            }
        ]
    },
    {
        id: 'Q_CS_02',
        scenario: "Someone keeps interrupting you during a conversation.",
        dimensionTargets: ['communication_style', 'conflict_posture'],
        options: [
            {
                label: 'I calmly point out that I\'d like to finish my thought',
                impacts: { communication_style: 0.2, conflict_posture: 0.15 }
            },
            {
                label: 'I wait for a pause and try to get my point across',
                impacts: { communication_style: -0.1, conflict_posture: -0.1 }
            },
            {
                label: 'I let it go but feel frustrated inside',
                impacts: { communication_style: -0.2, conflict_posture: -0.2 }
            }
        ]
    },
    {
        id: 'Q_CS_03',
        scenario: 'You need to explain a problem to someone who might not want to hear it.',
        dimensionTargets: ['communication_style'],
        options: [
            {
                label: 'I present the facts directly but with kindness',
                impacts: { communication_style: 0.2 }
            },
            {
                label: 'I ease into it with context before getting to the point',
                impacts: { communication_style: -0.1 }
            },
            {
                label: 'I hint at it and see if they pick up on it',
                impacts: { communication_style: -0.25 }
            }
        ]
    },
    {
        id: 'Q_CS_04',
        scenario: 'During a group discussion, you disagree with the majority opinion.',
        dimensionTargets: ['communication_style', 'conflict_posture', 'energy_orientation'],
        options: [
            {
                label: 'I share my perspective openly and explain my reasoning',
                impacts: { communication_style: 0.2, conflict_posture: 0.15, energy_orientation: 0.1 }
            },
            {
                label: 'I ask questions that might lead others to see my point',
                impacts: { communication_style: -0.1, conflict_posture: 0.0, energy_orientation: 0.0 }
            },
            {
                label: 'I keep my thoughts to myself unless asked directly',
                impacts: { communication_style: -0.2, conflict_posture: -0.15, energy_orientation: -0.15 }
            }
        ]
    },
    {
        id: 'Q_CS_05',
        scenario: "You're asked how you feel about a sensitive topic.",
        dimensionTargets: ['communication_style', 'emotional_availability'],
        options: [
            {
                label: 'I share my honest feelings clearly',
                impacts: { communication_style: 0.2, emotional_availability: 0.2 }
            },
            {
                label: 'I share some of what I feel, but keep parts private',
                impacts: { communication_style: 0.0, emotional_availability: 0.0 }
            },
            {
                label: "I deflect or give a vague answer",
                impacts: { communication_style: -0.2, emotional_availability: -0.2 }
            }
        ]
    },
    {
        id: 'Q_CS_06',
        scenario: 'You need to give someone directions to complete a task.',
        dimensionTargets: ['communication_style', 'decision_pace'],
        options: [
            {
                label: 'I give clear, step-by-step instructions',
                impacts: { communication_style: 0.2, decision_pace: 0.1 }
            },
            {
                label: 'I explain the goal and let them figure out the approach',
                impacts: { communication_style: -0.1, decision_pace: -0.1 }
            },
            {
                label: 'I show them by doing it together first',
                impacts: { communication_style: 0.0, decision_pace: -0.05 }
            }
        ]
    },

    // ==========================================================================
    // EMOTIONAL AVAILABILITY (6 questions)
    // ==========================================================================
    // Measures openness to emotional connection and vulnerability.
    // Higher scores = more emotionally open and accessible
    // Lower scores = more guarded, takes time to open up

    {
        id: 'Q_EA_01',
        scenario: 'A friend reaches out when they\'re going through a tough time.',
        dimensionTargets: ['emotional_availability', 'energy_orientation'],
        options: [
            {
                label: "I make time immediately and listen to what they need",
                impacts: { emotional_availability: 0.25, energy_orientation: 0.1 }
            },
            {
                label: 'I check in with them and offer practical support',
                impacts: { emotional_availability: 0.1, energy_orientation: 0.0 }
            },
            {
                label: 'I want to help but find emotional situations draining',
                impacts: { emotional_availability: -0.15, energy_orientation: -0.1 }
            }
        ]
    },
    {
        id: 'Q_EA_02',
        scenario: "You've been feeling stressed lately.",
        dimensionTargets: ['emotional_availability', 'communication_style'],
        options: [
            {
                label: 'I open up to someone I trust about what\'s going on',
                impacts: { emotional_availability: 0.25, communication_style: 0.1 }
            },
            {
                label: 'I mention I\'m stressed but keep the details private',
                impacts: { emotional_availability: 0.0, communication_style: -0.1 }
            },
            {
                label: 'I prefer to work through it on my own',
                impacts: { emotional_availability: -0.2, communication_style: -0.1 }
            }
        ]
    },
    {
        id: 'Q_EA_03',
        scenario: 'Someone you recently met shares something personal with you.',
        dimensionTargets: ['emotional_availability'],
        options: [
            {
                label: 'I appreciate their openness and share something back',
                impacts: { emotional_availability: 0.2 }
            },
            {
                label: 'I listen attentively and respond supportively',
                impacts: { emotional_availability: 0.1 }
            },
            {
                label: 'I feel a bit uncomfortable and keep things surface-level',
                impacts: { emotional_availability: -0.2 }
            }
        ]
    },
    {
        id: 'Q_EA_04',
        scenario: 'A close relationship has felt distant lately.',
        dimensionTargets: ['emotional_availability', 'conflict_posture'],
        options: [
            {
                label: 'I reach out and express that I want to reconnect',
                impacts: { emotional_availability: 0.2, conflict_posture: 0.1 }
            },
            {
                label: 'I give it some time and see if things improve naturally',
                impacts: { emotional_availability: -0.1, conflict_posture: -0.1 }
            },
            {
                label: 'I pull back too, matching their energy',
                impacts: { emotional_availability: -0.2, conflict_posture: -0.15 }
            }
        ]
    },
    {
        id: 'Q_EA_05',
        scenario: 'You notice someone seems upset but hasn\'t said anything.',
        dimensionTargets: ['emotional_availability', 'communication_style'],
        options: [
            {
                label: 'I gently ask if everything is okay',
                impacts: { emotional_availability: 0.2, communication_style: 0.1 }
            },
            {
                label: 'I stay nearby in case they want to talk',
                impacts: { emotional_availability: 0.1, communication_style: -0.1 }
            },
            {
                label: 'I give them space unless they bring it up',
                impacts: { emotional_availability: -0.15, communication_style: -0.15 }
            }
        ]
    },
    {
        id: 'Q_EA_06',
        scenario: "It's been a while since you connected with old friends.",
        dimensionTargets: ['emotional_availability', 'energy_orientation'],
        options: [
            {
                label: 'I reach out to catch up and reconnect',
                impacts: { emotional_availability: 0.2, energy_orientation: 0.15 }
            },
            {
                label: 'I think about reaching out but rarely follow through',
                impacts: { emotional_availability: -0.1, energy_orientation: -0.1 }
            },
            {
                label: 'I\'m okay with relationships fading if not maintained',
                impacts: { emotional_availability: -0.2, energy_orientation: -0.1 }
            }
        ]
    },

    // ==========================================================================
    // CONSISTENCY STYLE (6 questions)
    // ==========================================================================
    // Measures preference for routine vs. spontaneity.
    // Higher scores = prefers routine, predictability, structure
    // Lower scores = prefers spontaneity, variety, flexibility

    {
        id: 'Q_CT_01',
        scenario: 'Your weekly routine gets disrupted by an unexpected opportunity.',
        dimensionTargets: ['consistency_style', 'decision_pace'],
        options: [
            {
                label: 'I adjust my plans and take the opportunity',
                impacts: { consistency_style: -0.2, decision_pace: 0.1 }
            },
            {
                label: 'I weigh the trade-offs before deciding',
                impacts: { consistency_style: 0.0, decision_pace: -0.1 }
            },
            {
                label: 'I prefer to stick with my original plans',
                impacts: { consistency_style: 0.25, decision_pace: 0.0 }
            }
        ]
    },
    {
        id: 'Q_CT_02',
        scenario: 'You\'re planning a trip.',
        dimensionTargets: ['consistency_style'],
        options: [
            {
                label: 'I like having a detailed itinerary',
                impacts: { consistency_style: 0.25 }
            },
            {
                label: 'I plan the basics and leave room for spontaneity',
                impacts: { consistency_style: 0.0 }
            },
            {
                label: 'I prefer to figure things out as I go',
                impacts: { consistency_style: -0.25 }
            }
        ]
    },
    {
        id: 'Q_CT_03',
        scenario: 'Someone suggests trying a completely new restaurant last minute.',
        dimensionTargets: ['consistency_style', 'energy_orientation'],
        options: [
            {
                label: 'I\'m excited to try something new',
                impacts: { consistency_style: -0.2, energy_orientation: 0.1 }
            },
            {
                label: 'I\'m open to it if the reviews look good',
                impacts: { consistency_style: 0.0, energy_orientation: 0.0 }
            },
            {
                label: 'I\'d rather go somewhere I know I like',
                impacts: { consistency_style: 0.2, energy_orientation: -0.1 }
            }
        ]
    },
    {
        id: 'Q_CT_04',
        scenario: 'Your morning routine is interrupted by something unexpected.',
        dimensionTargets: ['consistency_style', 'emotional_regulation'],
        options: [
            {
                label: 'I adapt quickly and don\'t let it affect my day',
                impacts: { consistency_style: -0.15, emotional_regulation: 0.15 }
            },
            {
                label: 'I feel a bit off but manage',
                impacts: { consistency_style: 0.1, emotional_regulation: 0.0 }
            },
            {
                label: 'It throws off my whole day',
                impacts: { consistency_style: 0.2, emotional_regulation: -0.2 }
            }
        ]
    },
    {
        id: 'Q_CT_05',
        scenario: 'You have a free weekend with no plans.',
        dimensionTargets: ['consistency_style', 'energy_orientation'],
        options: [
            {
                label: 'I enjoy the freedom and see what happens',
                impacts: { consistency_style: -0.2, energy_orientation: 0.0 }
            },
            {
                label: 'I loosely plan a few activities',
                impacts: { consistency_style: 0.05, energy_orientation: 0.0 }
            },
            {
                label: 'I feel better when I have something scheduled',
                impacts: { consistency_style: 0.25, energy_orientation: 0.05 }
            }
        ]
    },
    {
        id: 'Q_CT_06',
        scenario: 'A project requires you to change your usual approach.',
        dimensionTargets: ['consistency_style', 'decision_pace'],
        options: [
            {
                label: 'I see it as a chance to learn something new',
                impacts: { consistency_style: -0.2, decision_pace: 0.0 }
            },
            {
                label: 'I adapt but prefer familiar methods',
                impacts: { consistency_style: 0.1, decision_pace: -0.05 }
            },
            {
                label: 'I find it challenging to change my workflow',
                impacts: { consistency_style: 0.2, decision_pace: -0.15 }
            }
        ]
    },

    // ==========================================================================
    // CONFLICT POSTURE (6 questions)
    // ==========================================================================
    // Measures approach to handling disagreements.
    // Higher scores = engages with conflict directly, addresses issues
    // Lower scores = avoids conflict, seeks harmony

    {
        id: 'Q_CP_01',
        scenario: 'You disagree with a decision made by your team.',
        dimensionTargets: ['conflict_posture', 'communication_style'],
        options: [
            {
                label: 'I voice my concerns and suggest alternatives',
                impacts: { conflict_posture: 0.25, communication_style: 0.15 }
            },
            {
                label: 'I share my view but go along with the group',
                impacts: { conflict_posture: 0.0, communication_style: 0.0 }
            },
            {
                label: 'I keep my disagreement to myself',
                impacts: { conflict_posture: -0.2, communication_style: -0.15 }
            }
        ]
    },
    {
        id: 'Q_CP_02',
        scenario: 'Someone makes a comment that bothers you.',
        dimensionTargets: ['conflict_posture', 'emotional_regulation'],
        options: [
            {
                label: 'I address it calmly and explain how I feel',
                impacts: { conflict_posture: 0.2, emotional_regulation: 0.15 }
            },
            {
                label: 'I let it pass but might bring it up later',
                impacts: { conflict_posture: -0.1, emotional_regulation: 0.0 }
            },
            {
                label: 'I try to forget about it to avoid tension',
                impacts: { conflict_posture: -0.25, emotional_regulation: -0.1 }
            }
        ]
    },
    {
        id: 'Q_CP_03',
        scenario: 'Two friends are in a disagreement and both come to you.',
        dimensionTargets: ['conflict_posture', 'emotional_availability'],
        options: [
            {
                label: 'I help mediate and encourage them to talk',
                impacts: { conflict_posture: 0.2, emotional_availability: 0.15 }
            },
            {
                label: 'I listen to both sides but stay neutral',
                impacts: { conflict_posture: 0.0, emotional_availability: 0.1 }
            },
            {
                label: 'I try to stay out of it',
                impacts: { conflict_posture: -0.2, emotional_availability: -0.1 }
            }
        ]
    },
    {
        id: 'Q_CP_04',
        scenario: 'You realize you were wrong about something in a discussion.',
        dimensionTargets: ['conflict_posture', 'communication_style'],
        options: [
            {
                label: 'I admit it openly and correct myself',
                impacts: { conflict_posture: 0.15, communication_style: 0.2 }
            },
            {
                label: 'I acknowledge it quietly and move on',
                impacts: { conflict_posture: 0.0, communication_style: -0.1 }
            },
            {
                label: 'I find it hard to admit when I\'m wrong',
                impacts: { conflict_posture: -0.15, communication_style: -0.15 }
            }
        ]
    },
    {
        id: 'Q_CP_05',
        scenario: 'Someone is consistently late to meetups.',
        dimensionTargets: ['conflict_posture'],
        options: [
            {
                label: 'I mention it directly and ask if we can adjust',
                impacts: { conflict_posture: 0.25 }
            },
            {
                label: 'I hint at it or plan differently without saying',
                impacts: { conflict_posture: -0.1 }
            },
            {
                label: "I don't say anything and just accept it",
                impacts: { conflict_posture: -0.25 }
            }
        ]
    },
    {
        id: 'Q_CP_06',
        scenario: 'A recurring issue keeps coming up in a relationship.',
        dimensionTargets: ['conflict_posture', 'emotional_availability'],
        options: [
            {
                label: 'I initiate a conversation to address it head-on',
                impacts: { conflict_posture: 0.2, emotional_availability: 0.15 }
            },
            {
                label: 'I wait for the right moment to bring it up',
                impacts: { conflict_posture: 0.0, emotional_availability: 0.0 }
            },
            {
                label: 'I hope it resolves on its own',
                impacts: { conflict_posture: -0.25, emotional_availability: -0.15 }
            }
        ]
    },

    // ==========================================================================
    // ENERGY ORIENTATION (5 questions)
    // ==========================================================================
    // Measures social energy preferences.
    // Higher scores = gains energy from social interaction
    // Lower scores = needs solitude to recharge

    {
        id: 'Q_EO_01',
        scenario: 'After a busy week, you finally have a free evening.',
        dimensionTargets: ['energy_orientation'],
        options: [
            {
                label: 'I reach out to friends to do something social',
                impacts: { energy_orientation: 0.25 }
            },
            {
                label: 'I might meet one person for a quiet catch-up',
                impacts: { energy_orientation: 0.0 }
            },
            {
                label: 'I prefer to spend the evening alone recharging',
                impacts: { energy_orientation: -0.25 }
            }
        ]
    },
    {
        id: 'Q_EO_02',
        scenario: 'You\'re at a gathering where you only know a few people.',
        dimensionTargets: ['energy_orientation', 'emotional_availability'],
        options: [
            {
                label: 'I enjoy meeting new people and starting conversations',
                impacts: { energy_orientation: 0.25, emotional_availability: 0.1 }
            },
            {
                label: 'I stick with who I know but am open if approached',
                impacts: { energy_orientation: 0.0, emotional_availability: 0.0 }
            },
            {
                label: 'I find it draining and look forward to leaving',
                impacts: { energy_orientation: -0.2, emotional_availability: -0.1 }
            }
        ]
    },
    {
        id: 'Q_EO_03',
        scenario: 'You have to work on a project.',
        dimensionTargets: ['energy_orientation', 'communication_style'],
        options: [
            {
                label: 'I prefer collaborating with others throughout',
                impacts: { energy_orientation: 0.2, communication_style: 0.1 }
            },
            {
                label: 'I like working alone but checking in with others',
                impacts: { energy_orientation: 0.0, communication_style: 0.0 }
            },
            {
                label: 'I do my best work alone with minimal interruption',
                impacts: { energy_orientation: -0.2, communication_style: -0.1 }
            }
        ]
    },
    {
        id: 'Q_EO_04',
        scenario: 'You have three social events in one weekend.',
        dimensionTargets: ['energy_orientation', 'consistency_style'],
        options: [
            {
                label: 'I look forward to a packed social weekend',
                impacts: { energy_orientation: 0.25, consistency_style: -0.1 }
            },
            {
                label: 'I\'ll go but will need some downtime between',
                impacts: { energy_orientation: 0.0, consistency_style: 0.1 }
            },
            {
                label: 'I\'d probably decline one or two to avoid burnout',
                impacts: { energy_orientation: -0.2, consistency_style: 0.15 }
            }
        ]
    },
    {
        id: 'Q_EO_05',
        scenario: 'A friend wants to introduce you to their friend group.',
        dimensionTargets: ['energy_orientation', 'emotional_availability'],
        options: [
            {
                label: 'I\'m excited to expand my social circle',
                impacts: { energy_orientation: 0.25, emotional_availability: 0.1 }
            },
            {
                label: "I'm open to it in the right setting",
                impacts: { energy_orientation: 0.0, emotional_availability: 0.0 }
            },
            {
                label: 'I prefer deepening existing friendships first',
                impacts: { energy_orientation: -0.15, emotional_availability: 0.1 }
            }
        ]
    },

    // ==========================================================================
    // DECISION PACE (5 questions)
    // ==========================================================================
    // Measures speed of decision-making.
    // Higher scores = decides quickly, trusts instincts
    // Lower scores = deliberate, thorough, needs time

    {
        id: 'Q_DP_01',
        scenario: 'You need to make an important decision with limited information.',
        dimensionTargets: ['decision_pace', 'emotional_regulation'],
        options: [
            {
                label: 'I trust my gut and make a call',
                impacts: { decision_pace: 0.25, emotional_regulation: 0.1 }
            },
            {
                label: 'I gather what I can quickly and decide',
                impacts: { decision_pace: 0.1, emotional_regulation: 0.0 }
            },
            {
                label: 'I delay until I have more information',
                impacts: { decision_pace: -0.25, emotional_regulation: -0.1 }
            }
        ]
    },
    {
        id: 'Q_DP_02',
        scenario: 'You\'re choosing between two equally good options.',
        dimensionTargets: ['decision_pace'],
        options: [
            {
                label: 'I pick one quickly and don\'t look back',
                impacts: { decision_pace: 0.25 }
            },
            {
                label: 'I take a bit of time but commit once I decide',
                impacts: { decision_pace: 0.0 }
            },
            {
                label: 'I go back and forth before finally choosing',
                impacts: { decision_pace: -0.25 }
            }
        ]
    },
    {
        id: 'Q_DP_03',
        scenario: 'Someone asks for your opinion on something right now.',
        dimensionTargets: ['decision_pace', 'communication_style'],
        options: [
            {
                label: 'I share my initial thoughts right away',
                impacts: { decision_pace: 0.2, communication_style: 0.15 }
            },
            {
                label: 'I give a quick answer but note it might change',
                impacts: { decision_pace: 0.1, communication_style: 0.0 }
            },
            {
                label: 'I prefer to think it over before responding',
                impacts: { decision_pace: -0.2, communication_style: -0.1 }
            }
        ]
    },
    {
        id: 'Q_DP_04',
        scenario: "You're shopping and can't decide between items.",
        dimensionTargets: ['decision_pace', 'consistency_style'],
        options: [
            {
                label: 'I pick one fairly quickly and move on',
                impacts: { decision_pace: 0.2, consistency_style: -0.1 }
            },
            {
                label: 'I take my time comparing before deciding',
                impacts: { decision_pace: -0.1, consistency_style: 0.1 }
            },
            {
                label: 'I might leave and come back another day',
                impacts: { decision_pace: -0.25, consistency_style: 0.15 }
            }
        ]
    },
    {
        id: 'Q_DP_05',
        scenario: 'A friend asks if you want to join them tomorrow for an activity.',
        dimensionTargets: ['decision_pace', 'energy_orientation', 'consistency_style'],
        options: [
            {
                label: 'I say yes or no right away based on how I feel',
                impacts: { decision_pace: 0.2, energy_orientation: 0.05, consistency_style: -0.1 }
            },
            {
                label: 'I check my schedule and get back to them soon',
                impacts: { decision_pace: 0.0, energy_orientation: 0.0, consistency_style: 0.1 }
            },
            {
                label: 'I need time to think about whether I have the energy',
                impacts: { decision_pace: -0.2, energy_orientation: -0.15, consistency_style: 0.0 }
            }
        ]
    },

    // ==========================================================================
    // CROSS-DIMENSIONAL QUESTIONS (Additional questions for depth)
    // ==========================================================================

    {
        id: 'Q_MX_01',
        scenario: 'You receive an invitation to an event happening tonight.',
        dimensionTargets: ['consistency_style', 'energy_orientation', 'decision_pace'],
        options: [
            {
                label: 'I love spontaneous plans and go for it',
                impacts: { consistency_style: -0.2, energy_orientation: 0.2, decision_pace: 0.2 }
            },
            {
                label: 'Depends on what I have going on',
                impacts: { consistency_style: 0.0, energy_orientation: 0.0, decision_pace: 0.0 }
            },
            {
                label: 'I prefer more notice before committing',
                impacts: { consistency_style: 0.2, energy_orientation: -0.1, decision_pace: -0.15 }
            }
        ]
    },
    {
        id: 'Q_MX_02',
        scenario: 'Someone shares a personal struggle with you unexpectedly.',
        dimensionTargets: ['emotional_availability', 'emotional_regulation', 'communication_style'],
        options: [
            {
                label: 'I listen fully and offer emotional support',
                impacts: { emotional_availability: 0.2, emotional_regulation: 0.1, communication_style: 0.1 }
            },
            {
                label: 'I try to help by offering advice or solutions',
                impacts: { emotional_availability: 0.0, emotional_regulation: 0.0, communication_style: 0.15 }
            },
            {
                label: 'I feel unsure how to respond',
                impacts: { emotional_availability: -0.15, emotional_regulation: -0.1, communication_style: -0.15 }
            }
        ]
    },
    {
        id: 'Q_MX_03',
        scenario: 'You strongly disagree with how something is being handled.',
        dimensionTargets: ['conflict_posture', 'communication_style', 'emotional_regulation'],
        options: [
            {
                label: 'I express my concerns clearly and directly',
                impacts: { conflict_posture: 0.2, communication_style: 0.2, emotional_regulation: 0.1 }
            },
            {
                label: 'I find a diplomatic way to share my view',
                impacts: { conflict_posture: 0.0, communication_style: 0.0, emotional_regulation: 0.1 }
            },
            {
                label: 'I keep it to myself to avoid rocking the boat',
                impacts: { conflict_posture: -0.25, communication_style: -0.2, emotional_regulation: -0.05 }
            }
        ]
    },
    {
        id: 'Q_MX_04',
        scenario: "You're feeling overwhelmed by responsibilities.",
        dimensionTargets: ['emotional_regulation', 'emotional_availability', 'communication_style'],
        options: [
            {
                label: 'I reach out to someone for support',
                impacts: { emotional_regulation: 0.1, emotional_availability: 0.2, communication_style: 0.1 }
            },
            {
                label: 'I take a break and recenter myself',
                impacts: { emotional_regulation: 0.15, emotional_availability: -0.1, communication_style: 0.0 }
            },
            {
                label: 'I push through even when struggling',
                impacts: { emotional_regulation: -0.15, emotional_availability: -0.2, communication_style: -0.1 }
            }
        ]
    },
    {
        id: 'Q_MX_05',
        scenario: 'You\'re asked to lead a group activity.',
        dimensionTargets: ['energy_orientation', 'communication_style', 'decision_pace'],
        options: [
            {
                label: 'I enjoy taking the lead and organizing',
                impacts: { energy_orientation: 0.2, communication_style: 0.15, decision_pace: 0.1 }
            },
            {
                label: "I'll do it if needed but prefer shared leadership",
                impacts: { energy_orientation: 0.0, communication_style: 0.0, decision_pace: 0.0 }
            },
            {
                label: "I'd rather let someone else take the lead",
                impacts: { energy_orientation: -0.15, communication_style: -0.1, decision_pace: -0.1 }
            }
        ]
    },
    {
        id: 'Q_MX_06',
        scenario: 'Plans change at the last minute due to circumstances.',
        dimensionTargets: ['consistency_style', 'emotional_regulation', 'decision_pace'],
        options: [
            {
                label: 'I adapt easily and make new plans',
                impacts: { consistency_style: -0.2, emotional_regulation: 0.2, decision_pace: 0.15 }
            },
            {
                label: "I'm a bit frustrated but go with it",
                impacts: { consistency_style: 0.1, emotional_regulation: 0.0, decision_pace: 0.0 }
            },
            {
                label: 'It really bothers me when plans change',
                impacts: { consistency_style: 0.2, emotional_regulation: -0.2, decision_pace: -0.1 }
            }
        ]
    }
];

// ============================================================================
// Helper Functions
// ============================================================================

/**
 * Get all questions from the question bank
 */
export function getAllQuestions(): Question[] {
    return QUESTION_BANK;
}

/**
 * Get a question by its ID
 */
export function getQuestionById(id: string): Question | undefined {
    return QUESTION_BANK.find(q => q.id === id);
}

/**
 * Get questions that target a specific dimension
 */
export function getQuestionsByDimension(dimension: PersonalityDimension): Question[] {
    return QUESTION_BANK.filter(q => q.dimensionTargets.includes(dimension));
}

/**
 * Get the count of questions per dimension
 */
export function getQuestionCountByDimension(): Record<PersonalityDimension, number> {
    const counts: Record<string, number> = {
        emotional_regulation: 0,
        communication_style: 0,
        emotional_availability: 0,
        consistency_style: 0,
        conflict_posture: 0,
        energy_orientation: 0,
        decision_pace: 0,
    };

    for (const question of QUESTION_BANK) {
        for (const dimension of question.dimensionTargets) {
            counts[dimension]++;
        }
    }

    return counts as Record<PersonalityDimension, number>;
}
