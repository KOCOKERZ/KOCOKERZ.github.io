function PostSignUp(doa,ayat,latin,artinya){
    var myHeaders = new Headers();
    myHeaders.append("Login", "rollygantengsekali");
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
        "doa": doa,
        "ayat": ayat,
        "latin": latin,
        "artinya": artinya
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    fetch("https://doa-doa-api-ahmadramadhan.fly.dev/api", requestOptions)
        .then(response => response.text())
        .then(result => console.log(result))
        .catch(error => console.log('error', error));
}

function PushButton(){
    doa=document.getElementById("doa").value;
    ayat=document.getElementById("ayat").value;
    latin=document.getElementById("latin").value;
    artinya=document.getElementById("artinya").value;
    PostSignUp(username,email,password);
}