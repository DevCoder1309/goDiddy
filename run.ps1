$String = "Hello UDP World"
$EncodedData = [System.Text.Encoding]::UTF8.GetBytes($String)
$UdpClient = New-Object System.Net.Sockets.UDPClient
$UdpClient.Connect("127.0.0.1", 8080)
$UdpClient.Send($EncodedData, $EncodedData.Length)
$UdpClient.Close()