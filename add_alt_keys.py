import os,string
fc=open("uiterm/keys.go","rb").read()
marker="//##altkeys##\n"
fc=fc.split(marker,1)[0]
fc+=marker
fc+="\n"
lines=[i.strip() for i in fc.split("\n")]
keys=[i.split(" ")[0].split("\t")[0] for i in lines if i.startswith("Key")]
keys=[i for i in keys if (i.startswith("Key") or i.startswith("Mouse"))]
fc+="const(\n"
for i in keys:
 if i.startswith("Key"):
  prefix="Key"
 elif i.startswith("Mouse"):
  prefix="Mouse"
 else:
  raise "no prefix for %s" % (i,)
 key=i[len(prefix):]
 fc+="%s%s%s Key = %s%s + (1<<16)\n" % (prefix,"Alt",key,prefix,key,)
for i in string.letters[:26].upper():
 fc+="%s%s%s Key = %s + (1<<16)\n" % ("Key","Alt",i,hex(ord(i.lower())),)

fc+=")\n"
fh=open("uiterm/keys.go.tmp","wb")
fh.write(fc)
fh.close()
os.rename("uiterm/keys.go.tmp","uiterm/keys.go")

