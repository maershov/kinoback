

#ifndef TESTTITANSCLASS_CONTAINERMANAGER_H
#define TESTTITANSCLASS_CONTAINERMANAGER_H


//CM

struct Package{
    unsigned int file_id;
    unsigned int user_id;
    unsigned int task_id;
    unsigned int package_id;
    unsigned int solution_id;
};


class ContainerManager{
public:
    void create_container();
    void send_task(Package package);
    void get_status();
    void remove_container();
    void get_load();
};


class Container{
public:
    void start_test();
    void send_task();
    void set_cont_state(class StateContainer* s);

    int load;
private:
    class StateContainer* state;
    Package package;

};

class StateContainer{
    virtual void start_test(Container* c);
    virtual void send_test(Container* c);
    virtual void handle();
};

class IsTested : public StateContainer{
    void start_test(Container* c);
    void send_test(Container* c);
    void handle();
};

class NotTested : public StateContainer{
    void start_test(Container* c);
    void send_test(Container* c);
    void handle();
};


#endif //TESTTITANSCLASS_CONTAINERMANAGER_H