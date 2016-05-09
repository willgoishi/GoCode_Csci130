var entry = document.querySelector('#Username');
var output = document.querySelector('#message');

function Process()
{
    console.log('ENTRY: ', entry.value);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/username/check');
    xhr.send(entry.value);
    xhr.addEventListener('readystatechange', function(){
        if (xhr.readyState === 4 && xhr.status === 200) {
            var taken = xhr.responseText;
            console.log('TAKEN:', taken, '\n\n');
            if (taken == 'true') {
                output.textContent = 'Sorry, user name already taken!';
            } else {
                output.textContent = '';
            }
        }
    });
}

entry.addEventListener('input', inputProcess);
