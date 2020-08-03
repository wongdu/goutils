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
        throw exception("���ԣ���B�Ĺ��캯�����׳�һ���쳣");
        cout << "���� B ����!" << endl;
    };
    ~B() { cout << "���� B ����!" << endl; };
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
    b = new B();  // <����
    cout << "���� Tester ����!" << endl;
}

Tester::~Tester()
{
    cout << "���� Tester ���� 1111 !" << endl;
    delete a;
    delete b;
    cout << "���� Tester ����!" << endl;
}
/******************������**********************/
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
    delete tes; // ɾ��NULLָ���ǰ�ȫ��
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
//��������ṹ
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
    set<string> allNames;    //�������в��ظ���name

    allNames.insert(list->data.name);  //���ȣ��ѵ�һ��Ԫ��ֱ�Ӽӽ�ȥ����Ϊ��һ�����ñȽ�
    List pre = list; //����ǰһ��ָ��
    List cur = list->next; //��ǰָ��

    while (cur != NULL)
    {
        //�����set��û���ҵ���ǰԪ�ص�name����˵���ü�¼���ظ������ü�¼��name���뵽set��������һ��
        if (allNames.find(cur->data.name) == allNames.end())
        {
            allNames.insert(cur->data.name);
            pre = cur;
            cur = cur->next;
        }
        else
        {
            //�����set���ҵ���ǰԪ�ص�name����˵���ظ�����ɾ����Ԫ�أ�Ȼ�������һ��
            Node* temp = cur; //�ȱ��浱ǰ�ڵ�ָ��
            cur = cur->next;
            pre->next = cur; //�ı�ָ��ָ��ɾ����ǰ�ڵ�
            //�������ͷŵ�temp�ڵ�
            delete temp;
        }
    }
    printList(list); //�����ҵĲ��Դ���
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