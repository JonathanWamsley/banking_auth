# banking_auth

This banking auth service is connected to the [banking repo](https://github.com/JonathanWamsley/banking). 
This service:
- stores login credentials
- generates jwt tokens 
- verifies jwt tokens
- checks for role access level (either user or admin level)

A users database stores login credentials and role account for a customer account or admin account.