# :point_down: LDAP 

LDAP (Lightweight Directory Access Protocol) เป็น protocol ที่ใช้สำหรับค้นหาข้อมูลของพนักงาน 

## :mag: การใช้งาน

### :bulb: System.DirectoryServices.DirectoryEntry (.NET 6 Support only windows)

```c#
/* import */
using System.DirectoryServices;
using System.DirectoryServices.AccountManagement;

/* code */
var ldapPath = "LDAP://DOMAIN.CORP";
var username = "domainname\\username";
var password = "p@ssw0rd";

using (var connection = new DirectoryEntry(ldapPath, username, password))
{
    if (connection == null)
    {
        throw new ArgumentNullException(nameof(connection));
    }

    using (var searcher = new DirectorySearcher(connection))
    {
        try
        {
            var permissions = searcher.FindAll();
            var userNotfound = (permissions == null || permissions.Count == 0);

            if(userNotfound)
            {
                throw new Exception("user not found");
            }

            var userCanAccess = (permissions > 0);

            // do some thing...
        }
        catch (Exception ex)
        {
            throw new Exception(ex.Message, ex.InnerException);
        }
    }
}
```

### :bulb: Novell.Directory.Ldap

```c#
/* import */
using Novell.Directory.Ldap;

/* code */
var ldapAddress = "domainname.corp";
var username = "username";
var password = "p@ssw0rd";

using (var connection = new LdapConnection { SecureSocketLayer = false })
{
    if (connection == null)
    {
        throw new ArgumentNullException(nameof(connection));
    }

    try
    {
        connection.Connect(ldapAddress, LdapConnection.DefaultPort);
        connection.Bind($"{username}@{ldapAddress}", password);

        var userCanAccess = connection.Bound;

        // do some thing...
    }
    catch (Exception ex)
    {
        throw new Exception(ex.Message, ex.InnerException);
    }
}
```

### :bulb: System.DirectoryServices.AccountManagement.PrincipalContext (.NET 6 Support only windows)

```c#
/* mport */
using System.DirectoryServices;
using System.DirectoryServices.AccountManagement;

/* code */
var ldapAddress = "domainname.corp";
var username = "username";
var password = "p@ssw0rd";

using (var context = new PrincipalContext(ContextType.Domain, ldapAddress))
{
    if (context == null)
    {
        throw new ArgumentNullException(nameof(context));
    }

    var userCanAccess = context.ValidateCredentials(username, password);

    // do some thing...
}
```
