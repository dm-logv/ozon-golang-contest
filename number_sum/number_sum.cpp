#include <iostream>
#include <stack>
#include <string>

using namespace std;

string DELIM = " ";
char   ZERO_CODE = 48;


void string_to_stack(string &s, stack<char> &digits) {
  for(string::size_type i = 0; i < s.size(); i++){
    digits.push(s[i] - ZERO_CODE);
  }
}


int main() {
  string line;
  getline(cin, line);
  cout << line << endl;

  size_t delim_position = line.find(DELIM);
  cout << delim_position << endl;

  string a_str = line.substr(0, delim_position);
  string b_str = line.substr(delim_position + 1);

  cout << a_str << endl;
  cout << b_str << endl;

  stack<char> a_digits, b_digits, result;
  
  string_to_stack(a_str, a_digits);
  string_to_stack(b_str, b_digits);

  int max_len = max(a_digits.size(), b_digits.size());
  char overflow = 0;
  char remainder = 0;
  
  cout << max_len << endl;

  for (int i = 0; i < max_len; i++) {
    char ax = a_digits.top(); a_digits.pop();
    char bx = b_digits.top(); b_digits.pop();

    char sum = ax + bx + overflow;
    overflow = sum / 10;
    remainder = sum % 10;

    result.push(remainder);
  }

  for (int i = 0; i < result.size(); i++) {
    cout << result.top() << " ";
    result.pop();
  }
  
  return 0;
}
