#include <iostream>
#include <vector>

//using namespace std;

class User {
    public:
        User(){};

        User(string username) {
            this->username = username;
            this->balance = 0;
        }

        void GiveBalance(double balance) {
            this->balance += balance;
        }

        void TakeBalance(double balance) {
            this->balance -= balance;
        }

        void CheckBalance() {
            cout << "Your balance is " << balance << endl;
        }

        double GetBalance() {
            return balance;
        }

    private:
        string username;
        double balance;
};

bool Roll();

int main() {
    string username;
    cout << "Enter username" << endl;
    cin >> username;
    User *current_user = new User(username);
    int answ;
    do {
        cout << "(1) Play" << endl;
        cout << "(2) Check balance" << endl;
        cout << "(3) Give balance" << endl;
        cout << "(0) Exit" << endl;
        cin >> answ;
        switch(answ) {
            case 1:
                {double bet = 0;
                cout << "Bet: " << endl;
                cin >> bet;
                double balance = current_user->GetBalance();
                if (balance > 0) {
                    if (Roll()) {
                        current_user->GiveBalance(bet);
                        break;
                    }
                    current_user->TakeBalance(bet);
                    break;
                }
                cout << "Not enought balance" << endl;
                }break;
            case 2:
                {current_user->CheckBalance();}
                break;
            case 3:
                {int newBalance = 0;
                cout << "How much?" << endl;
                cin >> newBalance;
                current_user->GiveBalance(newBalance);
                }break;
            default:
                {

                }
                break;
        }
    } while(answ != 0);
    return 0;
}

bool Roll() {
    int unumber;
    int rnumber = 1 + rand() % 10;
    cout << "Try to predict number" << endl;
    cin >> unumber;
    if(rnumber == unumber) {
        cout << "Congratulation!!!" << endl;
        return true;
    } else {
        cout << "You lost!!!!" << endl;
        cout << "Random number is " << rnumber << endl;
        return false;
    }
    return false;
}