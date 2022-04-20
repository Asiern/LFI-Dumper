Dump files over Local File Inclusion

```

 ▄▄▌  ·▄▄▄▪    ·▄▄▄▄  ▄• ▄▌• ▌ ▄ ·.  ▄▄▄·▄▄▄ .▄▄▄
 ██•  ▐▄▄ ██   ██· ██ █▪██▌·██ ▐███▪▐█ ▄█▀▄.▀·▀▄ █·
 ██ ▪ █  ▪▐█·  ▐█▪ ▐█▌█▌▐█▌▐█ ▌▐▌▐█· ██▀·▐▀▀▪▄▐▀▀▄
 ▐█▌ ▄██ .▐█▌  ██. ██ ▐█▄█▌██ ██▌▐█▌▐█▪·•▐█▄▄▌▐█•█▌
 .▀▀▀ ▀▀▀ ▀▀▀  ▀▀▀▀▀•  ▀▀▀ ▀▀  █▪▀▀▀.▀    ▀▀▀ .▀  ▀

     https://github.com/Asiern/LFI-Dumper


Usage: ./lfidumper -e 'http://target.com/page=' -d dictionary.txt

Options:
         -e : Endpoint url. -e 'http://target.com/page='
         -o : Output directory. -o output.
              If not specified the output directory will be './out'.
         -l : Login url. -l 'http://target/login'
         -p : Login POST payload. -p 'username=admin&password=admin&Login=Login'
         -d : Dictionary
         -f : Filter response body. Get response until string first appearance.
         -h : Show this menu

```

## Authentication
If authentication is required, you can define the login url and the POST payload to authenticate.

```
./lfidumper -e "http://target/?page=" -d dictionary.txt -l "http://target/login" -p "username=admin&password=password123"
```


