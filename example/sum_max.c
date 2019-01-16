int sum(int a){
 int s = 0;
 for(int i =0;i<a;i++){
   s += i;
 }
 return s;
}


int max(int a,int b){
  if(a>b){
    return a;
  }else{
    return b;
  }
}