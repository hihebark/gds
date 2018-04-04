#!bin/bash

#grep -o '[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\:[0-9]\{1,4\}'

#curl --silent https://proxy-list.org/french/index.php | grep -o "Proxy(['a-zA-Z0-9\')].*"\

curl --silent https://proxy-list.org/french/index.php | grep -o "Proxy(['a-zA-Z0-9\')].*" | sed 's/.*Proxy(.\(.*\).).*/\1/;' > $HOME/proxy.lst
_listproxy="$HOME/proxy.lst"
_nbrline=$(grep -c "." ${_listproxy})
for _line in `cat ${_listproxy}`
	do
	echo -e ` echo -e ${_line} | base64 --decode` >> ${_listproxy}
done
echo "$(tail -n +$((${_nbrline}+1)) ${_listproxy})" > "${_listproxy}"
cat $_listproxy
