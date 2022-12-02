//leetcode.com/problems/longest-valid-parentheses/

var str = "()())";
var start = 0, temp = 0;
var started = false;

for (let i = 0; i < str.length; i++){
	if(str[i] === "(" && started === false){
		start = i;
		started = true;
	}
	
	if(str[i] === ")"){
		temp = i;
	}

}

console.log(temp-start)
