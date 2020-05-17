Go script to use Route53 as dynamic DNS provider.

= Requirements =
* Go
* Route53 account with setup zone/domain name

= Usage =

Create file with AWS secrets somewhere
ex:
{
    "access_key" : "your access key",
    "secret_key" : "your secret key"
}

./routemaster -hosted-zone="example.com." -secrets-file=/Users/karl/.aws_secret -subdomain="subdomain"

This implementation is loosely based off https://github.com/dreamins/Route53DDNS-ruby
This is my first Go lang program so I wouldnt trust it with anything important

= TODO =

Refactor main method
Add tests
Support zone without subdomain
Better Error Handling
Add 2nd Host Grabbing IP Address

= Contributing =

Fork it
Create your feature branch (git checkout -b my-new-feature)
Commit your changes (git commit -am 'Added some feature')
Push to the branch (git push origin my-new-feature)
Create new Pull Request

= License =

Please see the included LICENSE.txt file.
