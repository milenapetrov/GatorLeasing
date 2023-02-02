for %%G in (*.sql) do sqlcmd /S servername /d releasedb -U username -P 
password -i"%%G"