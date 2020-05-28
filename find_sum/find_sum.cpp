#include <fstream>
#include <map>
#include <iostream>

using namespace std;


int main() {
  ifstream fin("input.txt");

  // Read a target
  int target;
  fin >> target;

  // cout << target << endl;
  
  // Read numbers
  map<int, int> numbers;
  int current;
  while (fin >> current) {
    numbers[current]++;
  }

  /*
  for(map<int, int>::iterator map_i = numbers.begin();
      map_i != numbers.end(); map_i++) {

    cout << map_i->first << ": " << map_i->second << endl;
  }
  */
  
  // Find pairs
  short result = 0;
  for(map<int, int>::iterator map_i = numbers.begin();
      map_i != numbers.end(); map_i++) {

    if (map_i->second == 0) { continue; }
    
    int a = map_i->first;
    int b = target - a;

    // cout << "a: " << a << " b: " << b << endl;

    if ((a == b && numbers[b] > 1)
        || (a != b && numbers[b] > 0)) {
      result = 1;
      break;
    }
  }

  // cout << result << endl; 
  ofstream fout("output.txt");
  fout << result;
  
  return 0;
}
