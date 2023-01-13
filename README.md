

<div align="center">

```
⚠️ NOTE: PP will migrate to use minikeyvalue aka https://github.com/geohot/minikeyvalue in v2
```

  <img width="128" alt="PP" src="https://user-images.githubusercontent.com/43297314/212246395-badef135-ad50-47ef-9baa-cfa5ed7dae7b.png">
  
  # PiedPiper     


_Acting as a **Middle** so that your **Out** can communicate with your **In**_

</div>



### 😠 Problem

- I have a home server and several Internet-connected nodemcu devices at different locations.
- I want to be able to access data collected by these devices from anywhere in the world using my Phone or Watch.

### 🤓 Techincal Problem

0. _Access the internal data.._
1. Without having to open a port in my home router.
2. I can't open ports on other locations.
3. Without having to use ssh tunneling or reverse proxy on home router.
4. I can't use ssh tunneling or reverse proxy on other devices.
5. Without having to setup a static IP.
6. Without turning my client device into a server and somehow connecting my Devices to client directly (also, see 2, 4)
7. Without trying to bypass firewall.

### 💡 Solution

- Send a POST request with key to a remote server to add / update data.

```sh
$ curl -X POST yourserver.url/super-secret-key -d 'I DID DONE IT... 20MB worth JSON b64 encoded data...'
```

- Send a GET request with the same key to get the data.

```sh
$ curl -X GET yourserver.url/super-secret-key
$ I DID DONE IT... 20MB worth JSON b64 encoded data...
```

![Freeform Thanku](https://user-images.githubusercontent.com/43297314/209245713-39635b00-8930-4f90-b4d4-cc3a756a03e3.png)


---

### ✨ Features

- Data persistence
- Rate Limiting
- Monitor API + Server status
- REST API for adding, updating and getting the data


✅ Deployed and tested 

![CleanShot 2022-12-22 at 19 52 12](https://user-images.githubusercontent.com/43297314/209249586-f5cbf301-baca-4c99-9038-0e7a00bfc7e0.png)

### Used by
[Watch Torrent](https://github.com/Aayush9029/WatchTorrent)
