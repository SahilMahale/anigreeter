#!/usr/bin/env perl
use strict;
use warnings;
use Getopt::Long;

my @quotes = (
    {
        quote => "I'll take a potato chip... and eat it!",
        character => "Light Yagami",
        anime => "Death Note"
    },
    {
        quote => "Whatever you lose, you'll find it again. But what you throw away you'll never get back.",
        character => "Himura Kenshin",
        anime => "Rurouni Kenshin"
    },
    {
        quote => "The world isn't perfect. But it's there for us, doing the best it can.",
        character => "Roy Mustang",
        anime => "Fullmetal Alchemist"
    },
    {
        quote => "If you don't like your destiny, don't accept it.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "People's lives don't end when they die. It ends when they lose faith.",
        character => "Itachi Uchiha",
        anime => "Naruto"
    },
    {
        quote => "If you don't take risks, you can't create a future!",
        character => "Monkey D. Luffy",
        anime => "One Piece"
    },
    # Naruto Uzumaki Quotes
    {
        quote => "I am strong because I have people to protect.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Failing doesn’t give you a reason to give up as long as you believe.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Don’t underestimate me! I don’t quit and I don’t run! You can act tough all you want! You’re not gonna scare me off! No way! I don’t care if I do get stuck as a Genin for the rest of my life! I’ll still be Hokage someday.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "There is no telling what kind of pain will come after me, but if I stop believing because of that, if the hero should change, it’ll turn into another story from the one my master left behind. Then it won’t be Naruto!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I already have many people who are important to me now, and I can’t let you hurt any of them. Even if I have to kill you!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "People who continue to put their lives on the line to defend their faith become heroes and continue to exist on in legend.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Because they saved me from myself, they rescued me from my loneliness. They were the first to accept me for who I am. They’re my friends.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "If he rips my arms off, I’ll kick him to death. If he rips my legs off, I’ll bite him to death. If he rips my head off, I’ll stare at him to death. And if he gouges out my eyes, I’ll curse him from beyond the grave!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "If I become as strong as you, will I really become like you? To die as a tool, that’s just too sad.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I won’t run away anymore. I won’t go back on my word. That is my ninja way!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "The pain of being alone is not an easy one to bear. Why is it that I can understand your pain?",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I can’t write novels like my master did. That’s why the sequel has to come from the life I live. No matter how great the pain, I’ll continue walking because that’s who Naruto is!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "If the father’s responsibility is to protect the child, then my responsibility is to exceed the father.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "To become Hokage is my dream!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Once you question your own belief, it’s over.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "So shut up about destiny and how people can’t change! Because unlike me, you’re not a failure.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "A true hero always arrives late.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "The only thing that can keep a fire from dying and give it more power, is wind.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Hard work is worthless for those that don’t believe in themselves.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I’m the only one who can bear the full brunt of your hate! It’s my job, no one else’s! I’ll bear the burden of your hatred, and we’ll die together!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "The many lives lost during long years of conflict—because of those selfless sacrifices, we are able to bathe in peace and prosperity now.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "It’s not the face that makes someone a monster, it’s the choices they make with their lives.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Sasuke! Sasuke! Sasuke! What does he have that I don’t have?",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Listen to yourself whining and complaining like some sorry little victim. You can whimper all day for all I care, you’re nothing but a coward!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Try it, trash! I’ll return the pain a thousand times over!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "My power isn’t the Rasengan, or Sage Jutsu, or the Nine Tails chakra. My power is right here.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Never give up.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Somebody told me I’m a failure. I’ll prove them wrong.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Before I became a ninja, I was a nobody; but I never gave up.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "If you don’t like your destiny, don’t accept it. Instead, have the courage to change it the way you want it to be!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I will become a bigger badass than my father, and a stronger shinobi than my mother!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "When people have different opinions, you should apply the majority vote.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "To ingrain this history within the new generation will be a vital cog in helping to maintain the peace.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "My name is Uzumaki Naruto. I like ramen. I hate the three minutes you have to wait while the water boils. And my dream is to one day become a Hokage. Then everyone will have to respect me at last.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "While you’re alive, you need a reason for your existence. Being unable to find one is the same as being dead.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I don’t understand what’s going on, but I’ll just pretend that I do.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "If we fight again, we’re both gonna die.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "If you don’t like the hand that fate’s dealt you with, fight for a new one!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "When you give up, your dreams and everything else, they’re gone.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "When people are protecting something truly special to them, they truly can become as strong as they can be.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Hey, you stupid fox! You’re in my body and you owe me rent. So for payment, I’ll take your chakra. You got that?",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I always wanted to be like you. I aspired to be like you. Because of that, I was glad that you wanted to fight me.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "The pain of being alone is completely out of this world, isn’t it? I don’t know why, but I understand your feelings so much, it actually hurts.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "We’re just some normal ninja, we don’t want to see what’s under your mask.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "If the world ended, who would you spend your last day with?",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I care more about others than I do myself, and I won’t let anyone hurt them.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "That’s why we endure. We are ninjas. I will never forget.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "Love breeds sacrifice, which in turn breeds hatred. Then, you can know pain.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I couldn’t understand what a parent’s love was like because you guys were never there, so I could only guess. But now I know, I live because you and Dad gave your lives for me and filled me up with love before the Nine-tails was inside me! So here I am, happy and healthy! I’m glad I ended up being your son!",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    },
    {
        quote => "I want to be with you. From now on, I want to spend all and every single one of my days until I die with you, and only you.",
        character => "Naruto Uzumaki",
        anime => "Naruto"
    }
);

# Command line options
my $anime;
my $character;
GetOptions(
    "anime|a=s" => \$anime,
    "character|c=s" => \$character
);

# Filter quotes based on arguments
my @filtered_quotes = @quotes;

if ($anime) {
    @filtered_quotes = grep {
        $_->{anime} =~ /$anime/i
    } @filtered_quotes;
}

if ($character) {
    @filtered_quotes = grep {
        $_->{character} =~ /$character/i
    } @filtered_quotes;
}

if (@filtered_quotes == 0) {
    die "No quotes found matching your criteria.\n";
}

# Select random quote
my $quote = $filtered_quotes[rand @filtered_quotes];

# Print formatted quote
print "\"$quote->{quote}\"\n";
print "  - $quote->{character} ($quote->{anime})\n";
