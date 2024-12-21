package quotes

type Quote struct {
	Text      string
	Character string
	Anime     string
}

func GetQuotes() []Quote {
	quotes := []Quote{
		{Text: "I am strong because I have people to protect.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Failing doesn’t give you a reason to give up as long as you believe.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Don’t underestimate me! I don’t quit and I don’t run! You can act tough all you want! You’re not gonna scare me off! No way! I don’t care if I do get stuck as a Genin for the rest of my life! I’ll still be Hokage someday.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "There is no telling what kind of pain will come after me, but if I stop believing because of that, if the hero should change, it’ll turn into another story from the one my master left behind. Then it won’t be Naruto!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I already have many people who are important to me now, and I can’t let you hurt any of them. Even if I have to kill you!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "People who continue to put their lives on the line to defend their faith become heroes and continue to exist on in legend.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Because they saved me from myself, they rescued me from my loneliness. They were the first to accept me for who I am. They’re my friends.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "If he rips my arms off, I’ll kick him to death. If he rips my legs off, I’ll bite him to death. If he rips my head off, I’ll stare at him to death. And if he gouges out my eyes, I’ll curse him from beyond the grave!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "If I become as strong as you, will I really become like you? To die as a tool, that’s just too sad.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I won’t run away anymore. I won’t go back on my word. That is my ninja way!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "The pain of being alone is not an easy one to bear. Why is it that I can understand your pain?", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I can’t write novels like my master did. That’s why the sequel has to come from the life I live. No matter how great the pain, I’ll continue walking because that’s who Naruto is!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "If the father’s responsibility is to protect the child, then my responsibility is to exceed the father.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "To become Hokage is my dream!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Once you question your own belief, it’s over.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "So shut up about destiny and how people can’t change! Because unlike me, you’re not a failure.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "A true hero always arrives late.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "The only thing that can keep a fire from dying and give it more power, is wind.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Hard work is worthless for those that don’t believe in themselves.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I’m the only one who can bear the full brunt of your hate! It’s my job, no one else’s! I’ll bear the burden of your hatred, and we’ll die together!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "The many lives lost during long years of conflict—because of those selfless sacrifices, we are able to bathe in peace and prosperity now.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "It’s not the face that makes someone a monster, it’s the choices they make with their lives.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Sasuke! Sasuke! Sasuke! What does he have that I don’t have?", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Listen to yourself whining and complaining like some sorry little victim. You can whimper all day for all I care, you’re nothing but a coward!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Try it, trash! I’ll return the pain a thousand times over!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "My power isn’t the Rasengan, or Sage Jutsu, or the Nine Tails chakra. My power is right here.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Never give up.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Somebody told me I’m a failure. I’ll prove them wrong.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Before I became a ninja, I was a nobody; but I never gave up.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "If you don’t like your destiny, don’t accept it. Instead, have the courage to change it the way you want it to be!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I will become a bigger badass than my father, and a stronger shinobi than my mother!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "When people have different opinions, you should apply the majority vote.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "To ingrain this history within the new generation will be a vital cog in helping to maintain the peace.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "My name is Uzumaki Naruto. I like ramen. I hate the three minutes you have to wait while the water boils. And my dream is to one day become a Hokage. Then everyone will have to respect me at last.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "While you’re alive, you need a reason for your existence. Being unable to find one is the same as being dead.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I don’t understand what’s going on, but I’ll just pretend that I do.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "If we fight again, we’re both gonna die.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "If you don’t like the hand that fate’s dealt you with, fight for a new one!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "When you give up, your dreams and everything else, they’re gone.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "When people are protecting something truly special to them, they truly can become as strong as they can be.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Hey, you stupid fox! You’re in my body and you owe me rent. So for payment, I’ll take your chakra. You got that?", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I always wanted to be like you. I aspired to be like you. Because of that, I was glad that you wanted to fight me.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "The pain of being alone is completely out of this world, isn’t it? I don’t know why, but I understand your feelings so much, it actually hurts.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "We’re just some normal ninja, we don’t want to see what’s under your mask.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "If the world ended, who would you spend your last day with?", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I care more about others than I do myself, and I won’t let anyone hurt them.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "That’s why we endure. We are ninjas. I will never forget.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "Love breeds sacrifice, which in turn breeds hatred. Then, you can know pain.", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I couldn’t understand what a parent’s love was like because you guys were never there, so I could only guess. But now I know, I live because you and Dad gave your lives for me and filled me up with love before the Nine-tails was inside me! So here I am, happy and healthy! I’m glad I ended up being your son!", Character: "Naruto Uzumaki", Anime: "Naruto"},
		{Text: "I want to be with you. From now on, I want to spend all and every single one of my days until I die with you, and only you.", Character: "Naruto Uzumaki", Anime: "Naruto"},
	}

	return quotes
}
