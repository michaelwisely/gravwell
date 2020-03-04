Dim fso, configFile
Dim params, paramsArray, CONFIGDIR
Dim CONFIG_LOG_LEVEL, CONFIG_CLEARTEXT_BACKEND_TARGET, CONFIG_INGEST_SECRET
Set fso = CreateObject("Scripting.FileSystemObject")

params = Session.Property("CustomActionData")
paramsArray = split(params, "|")
CONFIGDIR = paramsArray(0)
CONFIG_LOG_LEVEL = paramsArray(1)
CONFIG_CLEARTEXT_BACKEND_TARGET = paramsArray(2)

CONFIG_INGEST_SECRET = mid(params, len(CONFIGDIR) + len(CONFIG_LOG_LEVEL) + len(CONFIG_CLEARTEXT_BACKEND_TARGET) + 4)

Set configFile = fso.CreateTextFile(CONFIGDIR & "config.cfg", True)

configFile.WriteLine ("[Global]")
configFile.WriteLine ("Ingest-Secret = " & CONFIG_INGEST_SECRET)
configFile.WriteLine ("Connection-Timeout = 0")
configFile.WriteLine ("Verify-Remote-Certificates = true")
configFile.WriteLine ("#note that backslashes (\) are an escape character and must be escaped themselves")
configFile.WriteLine ("Cleartext-Backend-target=" & CONFIG_CLEARTEXT_BACKEND_TARGET & " #example of adding a cleartext connection")
configFile.WriteLine ("#Cleartext-Backend-target=127.1.0.1:4023 #example of adding another cleartext connection")
configFile.WriteLine ("#Encrypted-Backend-target=127.1.1.1:4024 #example of adding an encrypted connection")
configFile.WriteLine ("#Ingest-Cache-Path=""C:\\Program Files\\gravwell\\events.cache""")
configFile.WriteLine ("#Max-Ingest-Cache=1024 #Number of MB to store, localcache will only store 1GB before stopping.  This is a safety net")
configFile.WriteLine ("Log-Level=" & CONFIG_LOG_LEVEL)
configFile.WriteLine ()
configFile.WriteLine ("[EventChannel ""system""]")
configFile.WriteLine ("	#no Tag-Name means use the default tag")
configFile.WriteLine ("	Tag-Name=windows")
configFile.WriteLine ("	#no Provider means accept from all providers")
configFile.WriteLine ("	#no EventID means accept all event ids")
configFile.WriteLine ("	#no Level means pull all levels")
configFile.WriteLine ("	#no Max-Reachback means look for logs starting from now")
configFile.WriteLine ("	Channel=System #pull from the system channel")
configFile.WriteLine ("")
configFile.WriteLine ("[EventChannel ""application""]")
configFile.WriteLine ("	#no Tag-Name means use the default tag")
configFile.WriteLine ("	Tag-Name=windows")
configFile.WriteLine ("	#no Provider means accept from all providers")
configFile.WriteLine ("	#no EventID means accept all event ids")
configFile.WriteLine ("	#no Level means pull all levels")
configFile.WriteLine ("	#no Max-Reachback means look for logs starting from now")
configFile.WriteLine ("	Channel=Application #pull from the system channel")
configFile.WriteLine ()
configFile.WriteLine ("[EventChannel ""security""]")
configFile.WriteLine ("	#no Tag-Name means use the default tag")
configFile.WriteLine ("	Tag-Name=windows")
configFile.WriteLine ("	#no Provider means accept from all providers")
configFile.WriteLine ("	#no EventID means accept all event ids")
configFile.WriteLine ("	#no Level means pull all levels")
configFile.WriteLine ("	#no Max-Reachback means look for logs starting from now")
configFile.WriteLine ("	Channel=Security #pull from the system channel")
configFile.WriteLine ()
configFile.WriteLine ("[EventChannel ""setup""]")
configFile.WriteLine ("	#no Tag-Name means use the default tag")
configFile.WriteLine ("	Tag-Name=windows")
configFile.WriteLine ("	#no Provider means accept from all providers")
configFile.WriteLine ("	#no EventID means accept all event ids")
configFile.WriteLine ("	#no Level means pull all levels")
configFile.WriteLine ("	#no Max-Reachback means look for logs starting from now")
configFile.WriteLine ("	Channel=Setup #pull from the system channel")
configFile.WriteLine ()
configFile.WriteLine ("############# EXAMPLE additional listeners #############")
configFile.WriteLine ("#[EventChannel ""sysmon""]")
configFile.WriteLine ("#	Tag-Name=sysmon")
configFile.WriteLine ("#	Channel=""Microsoft-Windows-Sysmon/Operational""")
configFile.WriteLine ("#	Max-Reachback=24h  #reachback must be expressed in hours (h), minutes (m), or seconds(s)")
configFile.WriteLine ("#")
configFile.WriteLine ("#")
configFile.WriteLine ("#[EventChannel ""Application""]")
configFile.WriteLine ("#	Channel=Application #pull from the application channel")
configFile.WriteLine ("#	Tag-Name=winApp #Apply a new tag name")
configFile.WriteLine ("#	Provider=Windows System #Only look for the provider ""Windows System""")
configFile.WriteLine ("#	EventID=1000-4000 #Only look for event IDs 1000 through 4000")
configFile.WriteLine ("#	Level=verbose #Only look for verbose entries")
configFile.WriteLine ("#	Max-Reachback=72h #start looking for logs up to 72 hours in the past")
configFile.WriteLine ("#")
configFile.WriteLine ("#")
configFile.WriteLine ("#[EventChannel ""System Critical and Error""]")
configFile.WriteLine ("#	Channel=System #pull from the system channel")
configFile.WriteLine ("#	Tag-Name=winSysCrit #Apply a new tag name")
configFile.WriteLine ("#	#no provider, we want everything")
configFile.WriteLine ("#	#no eventID, give em all")
configFile.WriteLine ("#	Level=critical #look for critical entries")
configFile.WriteLine ("#	Level=error #AND for error entries")
configFile.WriteLine ("#	Max-Reachback=96h #start looking for logs up to 96 hours in the past")
configFile.WriteLine ("#")
configFile.WriteLine ("#")
configFile.WriteLine ("#[EventChannel ""Security prune""]")
configFile.WriteLine ("#	Channel=Security #pull from the security channel")
configFile.WriteLine ("#	Tag-Name=winSec #Apply a new tag name")
configFile.WriteLine ("#	#no provider, we want everything")
configFile.WriteLine ("#	#no level implies all levels")
configFile.WriteLine ("#	EventID=-400 #ignore event ID 400")
configFile.WriteLine ("#	EventID=-401 #AND ignore event ID 401")
configFile.WriteLine ("#")
configFile.WriteLine ("#")

configFile.Close
