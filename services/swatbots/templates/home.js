{{define "content"}}


function Loop(){

for (var i =0 ; i < 18; i ++ ) {
    rest(i);
}

for (var i =1 ; i < 18; i= i+3 ) {
    rest2(i);
}

}
//Loop();
window.setInterval(function(){Loop()},2000);

function rest2(i){
    
   // console.log(i.toString());
    $.ajax({
	url: 'http://api.swatbots.com/sensor/read?id=sp' + i.toString(),
	type: 'GET',
	datatype: 'json',
	data : "testing",
	success: function(json){ handle2(json , i); }

    }); 

}

function handle2(json, i) {
    //console.log("inside handle 2 ");
    //console.log(json);
    range = document.getElementsByTagName('input')[(i-1)/3];
    if (json.d.Value.length > 0) { 
   range.value = json.d.Value[0];
    }
}
function rest(i){
    
    $.ajax({
	url: 'http://api.swatbots.com/sensor/read?id=sp' + i.toString() ,
	type: 'GET',
	datatype: 'json',
	data : "testing",
	success: function(json){ handle(json , i); }

    }); 

}

function handle(json, i){
    span = document.getElementsByTagName('span')[i];

    if(i%3 == 0){    
    span.innerHTML = json.d.Value[0] + 'Lux'; 
    }
   
 if(i%3 == 1){    
    span.innerHTML = json.d.Value[0] + '%';
    }
 if(i%3 == 2){    
    span.innerHTML = json.d.Value[0]+ 'mAmp';
    }
}

{{end}}