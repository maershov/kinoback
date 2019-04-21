

#ifndef TESTTITANSCLASS_FRONTMANAGER_H
#define TESTTITANSCLASS_FRONTMANAGER_H

//FM


#include <iostream>
#include <string>

using std::string;

struct Human{
    char *name;
    char *surname;
    char *password;
    unsigned int username;
    unsigned int group_id;
};

struct Task{
    unsigned int task_id;
    char *description;
    unsigned int script_id;
    unsigned int contest_id;
};

struct Contest{
    unsigned int contest_id;
    unsigned int *human_list_id;
    unsigned int *tasks_list_id;
    //unsigned int admin_id;
    //unsigned int username;
};

struct Package{
    unsigned int file_id;
    unsigned int user_id;
    unsigned int task_id;
    unsigned int package_id;
    unsigned int solution_id;
};

struct Solution{
    unsigned int solution;
    int result;
};

struct StatusAP{
    unsigned int percent;
    unsigned int file_id;
};

struct Session{
    int session_id;
};

class ListenServer{
    Session listen(char port);
};

class RequestParser{
    int command_parser(string command);
};

class FrontManager{
public:
    void set_command(int command);
    void make_command(int command_value);
    int create_user(Human human);
    int create_contest(Contest contest);
    int create_task(Task task);
    int delete_user(Human human);
    int delete_contest(Contest contest);
    int delete_task(Task task);
    int start_contest(Contest contest);
    unsigned int* get_stat(unsigned int contest_id);
    Solution get_solution(unsigned int solution_id);
    unsigned int *get_list_competitors(unsigned int contest_id);

private:
    int command;
};

class IClientBD{
public:
    virtual void update() = 0;
};

class ClientBD : public IClientBD{
public:
    void send_human(Human human){
        std::cout << "hello from send_H\n";
    }
    void send_task(Task task){
        std::cout << "hello from send_T\n";
    }
    void send_contest(Contest contest){
        std::cout << "hello from send_C\n";
    }
    void send_package(Package package){
        std::cout << "hello from send_P\n";
    }

    Solution get_solution(unsigned int solution_id);
    unsigned int *get_list_competitors(unsigned int contest_id);
    void update() override{
        std::cout << "Hello from update CBD";
    };
};

class IClientScheduller{
public:
    virtual void update() = 0;
};

class ClientScheduller : public IClientScheduller{
public:
    void send_package(Package package);
    void update() override{
        std::cout << "Hello from update CS";
    };
};

class IClientStorage{
public:
    virtual void update() = 0;
};

class ClientStorage : public IClientStorage{
public:
    void send_to_storage(Package package);
    void update() override{
        std::cout << "Hello from update CSt";
    }
private:
    unsigned int file_id;
    uint8_t file;
    string description;

};

class IClientAP{
    virtual void update() = 0;
};

class ClientAP : public IClientAP{
    StatusAP get_status(unsigned int package_id);
    void update() override{
        std::cout << "Hello from CAP";
    }
};


#endif //TESTTITANSCLASS_FRONTMANAGER_H