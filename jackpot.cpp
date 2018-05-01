// C++: Use roll() to decide a certain set of numbers, display them under some suspense!
#include <cstdlib>
#include <iostream>
#include <ctime>
#include <set>
#include <unistd.h>

void roll(size_t row_length, size_t max) {
    std::set<int> winners;
    while (winners.size() != row_length) {
        winners.insert(std::rand()%max + 1);
        usleep((std::rand()%90+10)*1000);
    }
    for (auto winner : winners) {
        sleep(std::rand()%5+1);
        std::cout << winner << " " << std::flush;
    }
    std::cout << std::endl;
}

int main() {
    std::srand(std::time(0));
    std::cout << "Normal numbers:" << std::endl;
    roll(5, 50);

    std::cout << std::endl;
    std::cout << "Star numbers" << std::endl;
    roll(2, 10);
}
