#include <iostream>
//#include <string>

#include <grpcpp/grpcpp.h>
#include "service.grpc.pb.h"
#include "helloworld.grpc.pb.h"
#include "syncaliyunoss.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using service::AliyunOssRequest;
using service::AliyunOssResponse;
using service::AliyunOss;

using helloworld::HelloRequest;
using helloworld::HelloReply;
using helloworld::Greeter;

/*
using syncaliyunoss::AliyunOssRequest;
using syncaliyunoss::AliyunOssReply;
using syncaliyunoss::SyncAliyunOssFile;
*/

using namespace std;
 
#pragma once
#include <string>
#include <iostream>
using namespace std;

class A {
public:
    A() {};
};

class B {
public:
    B() {
        throw exception("测试：在B的构造函数中抛出一个异常");
        cout << "构造 B 对象!" << endl;
    };
    ~B() { cout << "销毁 B 对象!" << endl; };
};

class Tester {
public:
    Tester(const string& name, const string& address);
    ~Tester();
private:
    string theName;
    string theAddress;
    A aa;
    A* a;
    B* b;
};

Tester::Tester(const string& name, const string& address) :
    theName(name),
    theAddress(address)
{
    a = new A();
    b = new B();  // <——
    cout << "构造 Tester 对象!" << endl;
}

Tester::~Tester()
{
    cout << "销毁 Tester 对象 1111 !" << endl;
    delete a;
    delete b;
    cout << "销毁 Tester 对象!" << endl;
}
/******************测试类**********************/
int main6()
{
    Tester* tes = NULL;
    try
    {
        tes = new Tester("songlee", "201");
    }
    catch (exception e)
    {
        cout << e.what() << endl;
    };
    cout << "main!" << endl;
    delete tes; // 删除NULL指针是安全的
    getchar();
    return 0;
}

#include <set>
#include <iostream>
#include <string>
using namespace std;

struct Data {
    string name;
};
//定义链表结构
struct Node {
    Data data;
    Node* next;
};
typedef Node* List;
//For test
string names[] = { "hello", "world", "hello","hello world", "zht","zhengh", "zht", "hello", "world", "hello world" };
void createList(List& list)
{
    list = new Node;
    list->data.name = "zhtsuc";
    List p = list;
    for (int i = 0; i < 10; i++)
    {
        Node* node = new Node;
        node->data.name = names[i];
        p->next = node;
        p = node;
    }
    p->next = NULL;
}
void printList(List list)
{
    List p = list;
    while (p != NULL)
    {
        cout << p->data.name << endl;
        p = p->next;
    }
}
int main5(int argc, char* argv[])
{
    List list;
    createList(list);
    set<string> allNames;    //保存所有不重复的name

    allNames.insert(list->data.name);  //首先，把第一个元素直接加进去。因为第一个不用比较
    List pre = list; //保存前一个指针
    List cur = list->next; //当前指针

    while (cur != NULL)
    {
        //如果在set中没有找到当前元素的name，则说明该记录不重复，将该记录的name加入到set，继续下一个
        if (allNames.find(cur->data.name) == allNames.end())
        {
            allNames.insert(cur->data.name);
            pre = cur;
            cur = cur->next;
        }
        else
        {
            //如果在set中找到当前元素的name，则说明重复，即删掉此元素，然后继续下一个
            Node* temp = cur; //先保存当前节点指针
            cur = cur->next;
            pre->next = cur; //改变指针指向，删除当前节点
            //在这里释放掉temp节点
            delete temp;
        }
    }
    printList(list); //这是我的测试代码
    return 0;
}

int main() {
    std::string target_str = "localhost:7877";
    std::shared_ptr<Channel> channel = grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials());
    std::unique_ptr<syncaliyunoss::SyncAliyunOssFile::Stub> stub_ = syncaliyunoss::SyncAliyunOssFile::NewStub(channel);

    syncaliyunoss::AliyunOssRequest request;
    request.set_endpoint("oss-cn-shenzhen.aliyuncs.com");
    request.set_bucket_name("res-leimans-com-1");
    request.set_object_name_prefix("user_game_data/");
    request.set_file_name("10833_10111.7z");
    request.set_md5sum_value("e82c3a5afe9d41fac9aa54223b964b6e");
    //request.set_timestamp(1594795169);
    time_t now = 0;
    now = ::time(NULL);
    request.set_timestamp(now);
   // struct tm* tm = localtime_s(&now);

    syncaliyunoss::AliyunOssReply reply;
    ClientContext context;

    // The actual RPC.
    /*Status status = stub_->SyncOssFile(&context, request, &reply);
    if (status.ok()) {
        std::cout << "rpc call ok" << std::endl;
        std::cout << "rpc call ok" << std::endl;
    }
    else {
        std::cout << status.error_code() << ": " << status.error_message()
            << std::endl;
    }*/

    ::google::protobuf::Empty empty;
	Status status = stub_->SyncOssNow(&context, empty, &reply);
	if (status.ok()) {
		std::cout << "rpc call ok" << std::endl;
		std::cout << reply.message() << std::endl;
        std::cout << reply.errcode() << std::endl;
	}
	else {
		std::cout << status.error_code() << ": " << status.error_message()
			<< std::endl;
	}

    return 0;
}

int main3() {
    std::string target_str = "localhost:7877";    
    std::shared_ptr<Channel> channel = grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials());
    std::unique_ptr<syncaliyunoss::SyncAliyunOssFile::Stub> stub_ = syncaliyunoss::SyncAliyunOssFile::NewStub(channel);

    syncaliyunoss::AliyunOssRequest request;
    request.set_endpoint("oss-cn-shenzhen.aliyuncs.com");
    request.set_bucket_name("res-leimans-com-1");
    request.set_file_name("");
    request.set_md5sum_value("");
    request.set_timestamp(1594795169);

    syncaliyunoss::AliyunOssReply reply;
    ClientContext context;

    // The actual RPC.
    Status status = stub_->SyncOssFile(&context, request, &reply);
    if (status.ok()) {
        std::cout << "rpc call ok" << std::endl;
    }
    else {
        std::cout << status.error_code() << ": " << status.error_message()
            << std::endl;
    }

    return 0;
}

int main2() {
    std::string target_str = "localhost:7877";
    //std::string target_str = "0.0.0.0:7877";
    //std::string target_str = "127.0.0.1:7877";
    std::shared_ptr<Channel> channel = grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials());
    std::unique_ptr<Greeter::Stub> stub_ = Greeter::NewStub(channel);

    HelloRequest request;
    request.set_name("userName");

    HelloReply reply;
    ClientContext context;

    // The actual RPC.
    Status status = stub_->SayHello(&context, request, &reply);
    if (status.ok()) {
        std::cout << "rpc call ok" << std::endl;
    }
    else {
        std::cout << status.error_code() << ": " << status.error_message()
            << std::endl;
    }

    return 0;
}

int main1()
{
    std::string target_str = "localhost:7877";
    //std::string target_str = "0.0.0.0:7877";
    //std::string target_str = "127.0.0.1:7877";
    std::shared_ptr<Channel> channel= grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials());
    std::unique_ptr<AliyunOss::Stub> stub_ = AliyunOss::NewStub(channel);

    AliyunOssRequest request;
    /*request.set_endpoint("");
    request.set_bucket_name("");
    request.set_object_name_prefix("");
    request.set_file_name("");
    request.set_md5sum_value("");
    request.set_timestamp(1594795169);*/
    request.set_endpoint("");
    request.set_bucketname("");
    request.set_filename("");
    request.set_md5sumvalue("");
    request.set_timestamp(1594795169);

    AliyunOssResponse reply;
    ClientContext context;

    // The actual RPC.
    Status status = stub_->SyncOssFile(&context, request, &reply);
    if (status.ok()) {
        std::cout << "rpc call ok" << std::endl;
    }
    else {
        std::cout << status.error_code() << ": " << status.error_message()
            << std::endl;
    }

    return 0;
}