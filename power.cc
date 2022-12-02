#include<stdio.h>
#include<stdlib.h>
#include <string>
using namespace std;

int main(int argc, char **argv){
    if(argc == 1){
        return -1;
    }else if(argc == 2){
        printf("%d",argc);
    }

    
    int first = atoi(argv[1]);
    int second = atoi(argv[2]);
    int result = first;

    if(second == 0){
        result = 1;
    }else{
        for(int i = 1; i < second; i++){
            result = result * first;
        }
    }

    if(second < 0){
        printf("1/%d", result);
    }else{
        printf("%d", result);
    }
    return 0;
}
