# go-oracle-sample

1.install gopkg.in/goracle.v2
    
`＄go get gopkg.in/goracle.v2`

2.ref https://www.oracle.com/technetwork/topics/intel-macsoft-096467.html

3.Download the desired Instant Client ZIP files.

4.Add links to ~/lib or /usr/local/lib to enable applications to find the libraries.

`＄ln -s ~/instantclient_12_2/libclntsh.dylib /usr/local/lib`
	
`＄ln -s ~/instantclient_12_2/libclntsh.dylib.12.1 /usr/local/lib`

5.If you intend to co-locate optional Oracle configuration files such as tnsnames.ora, sqlnet.ora, ldap.ora, or oraaccess.xml
with Instant Client, then create a network/admin subdirectory.

`＄mkdir -p ~/instantclient_12_2/network/admin`
	
6.dbusername/dbpassword@serviceName

`db, err := sql.Open("goracle", fmt.Sprintf("%s/%s@%s", dbusername, dbpassword, serviceName))`

**Run GO app**

`＄go run app.go database.go -log_dir=./`

**Create a new repository on the command line**

`echo "# go-oracle-sample" >> README.md`

`git init`

`git add README.md`

`git commit -m "first commit"`

`git remote add origin https://github.com/posheng/go-oracle-sample.git`

`git push -u origin master`