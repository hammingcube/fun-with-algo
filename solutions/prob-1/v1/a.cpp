# include <iostream>
# include <fstream>

using namespace std;

int func(int a, int b) {
	return a+b;
}

int main() {
	int a, b, result;
	//cout << "Hello\n";	
	ifstream ifile("STDIN");
	while(ifile >> a) {
		ifile >> b;
		result = func(a,b);
		cout << result << endl;
	}
	return 0;
}
