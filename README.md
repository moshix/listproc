This is a bitnet, or better an HNET list processor for mailinglists
-------------------------------------------------------------------

Dec 2019 by moshix
Apache license




INSTALL
------

git clone
git build listproc.go


RUN
---

After building it run it with ./listproc


User guide:

LISTPROC Commands

LISTPROC /CREATE discussion
LISTPROC /SEARCH keywords keywords keywords returns list of discussion with they keywords
LISTPROC /RECEIVE discussion1 discussion2 discussion3 returns 3 file sends with those full discussions
LISTPROC /SUBSCRIBE   discussion
LISTPROC /UNSUBSCRIBE discussion
LISTPROC /LIST     sends a list of all discussions by file send

