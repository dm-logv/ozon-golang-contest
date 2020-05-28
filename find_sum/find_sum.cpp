#include <fstream>
#include <map>
#include <iostream>

using namespace std;


int main() {
  ifstream fin("input.txt");

  // Read a target
  int target;
  fin >> target;
  
  // Read numbers
  map<int, int> numbers;
  int current;
  while (fin >> current) {
    numbers[current]++;
  }
  
  // Find pairs
  short result = 0;
  for(map<int, int>::iterator map_i = numbers.begin();
      map_i != numbers.end(); map_i++) {

    if (map_i->second == 0) { continue; }
    
    int a = map_i->first;
    int b = target - a;

    if ((a == b && numbers[b] > 1)
        || (a != b && numbers[b] > 0)) {
      result = 1;
      break;
    }
  }

  ofstream fout("output.txt");
  fout << result;
  
  return 0;
}
