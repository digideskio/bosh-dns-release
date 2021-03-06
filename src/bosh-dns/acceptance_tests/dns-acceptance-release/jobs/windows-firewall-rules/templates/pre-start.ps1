try
{
    New-NetFirewallRule -DisplayName "bosh dns server TCP" -Direction Inbound -LocalPort 53 -Protocol TCP
    New-NetFirewallRule -DisplayName "bosh dns server UDP" -Direction Inbound -LocalPort 53 -Protocol UDP
}
catch
{
    $Host.UI.WriteErrorLine($_.Exception.Message)
    Exit 1
}
Exit 0
