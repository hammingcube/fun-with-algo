# include <iostream>

using namespace std;

int func(int a, int b) {
	return a+b;
}

int main() {
	int a, b, result;
	while(cin >> a) {
		cin >> b;
		result = func(a,b);
		cout << result << endl;
	}
	return 0;
}
