# Crystal: Use all numbers, shuffle and display the first ones
print (1..50).to_a.shuffle(Random::Secure)[0...5].sort.to_s + "\n"
print (1..10).to_a.shuffle(Random::Secure)[0...2].sort.to_s + "\n"
