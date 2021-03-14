#include <string>

using namespace std;

struct HelloRequest {
  string msg;
};

struct HelloResponse {
  string msg;
};

class HelloService {
public:
  static HelloResponse Hello(HelloRequest);
};

int main() { }

