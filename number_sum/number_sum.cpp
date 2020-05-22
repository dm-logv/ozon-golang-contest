#include <iostream>
#include <stack>
#include <string>
#include <utility>

using namespace std;

string DELIM = " ";
char   ZERO_CODE = 48;


pair<string, string> split_string(string &delim, string &line) {
  size_t delim_position = line.find(DELIM);

  return make_pair(
      line.substr(0, delim_position),
      line.substr(delim_position + 1));
}


stack<char>* string_to_stack(string s) {
  stack<char>* digits = new stack<char>;
  
  for(string::size_type i = 0; i < s.size(); i++){
    digits->push(s[i] - ZERO_CODE);
  }

  return digits;
}


char pop_or_zero(stack<char> &stack) {
  char value = 0;
  if (stack.size() > 0) {
    value = stack.top();
    stack.pop();
  }
  return value;
}


void print_stack(stack<char> &stack) {
  while (!stack.empty()) {
    cout << (short) stack.top();
    stack.pop();
  }
}


int main() {
  string line;
  getline(cin, line);

  pair<string, string> splitted = split_string(DELIM, line);
  string a_str = splitted.first;
  string b_str = splitted.second;
  
  stack<char>* a_digits = string_to_stack(a_str);
  stack<char>* b_digits = string_to_stack(b_str);

  int max_len = max(a_digits->size(), b_digits->size());

  stack<char> result;
  char overflow = 0;
  char remainder = 0;
  
  for (int i = 0; i < max_len; i++) {
    char ax = pop_or_zero(*a_digits);
    char bx = pop_or_zero(*b_digits);

    char sum = ax + bx + overflow;
    overflow = sum / 10;
    remainder = sum % 10;
    
    result.push(remainder);
  }

  if (overflow > 0) {
    result.push(overflow);
  }

  print_stack(result);
  
  return 0;
}

