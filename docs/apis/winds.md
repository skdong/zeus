# 接口

## 普通Restful接口

http://localhost/8080/winds

支持参数

- device_id

- limit

- start_at

- end_at

- start

- limit

Response
```json
{
  "Num": 5,
  "SpeedAvg": 2.74,
  "DirectionAvg": 229,
  "winds": [
    {
      "CreateAt": "2020-04-25T15:54:32Z",
      "DeviceId": "Q",
      "Direction": 229,
      "Id": 2549,
      "Speed": 2.74,
      "Unit": "M"
    },
    {
      "CreateAt": "2020-04-25T15:54:31Z",
      "DeviceId": "Q",
      "Direction": 229,
      "Id": 2548,
      "Speed": 2.74,
      "Unit": "M"
    },
    {
      "CreateAt": "2020-04-25T15:54:30Z",
      "DeviceId": "Q",
      "Direction": 229,
      "Id": 2547,
      "Speed": 2.74,
      "Unit": "M"
    },
    {
      "CreateAt": "2020-04-25T15:54:29Z",
      "DeviceId": "Q",
      "Direction": 229,
      "Id": 2546,
      "Speed": 2.74,
      "Unit": "M"
    },
    {
      "CreateAt": "2020-04-25T15:54:28Z",
      "DeviceId": "Q",
      "Direction": 229,
      "Id": 2545,
      "Speed": 2.74,
      "Unit": "M"
    }
  ]
}
```

## Websocket接口

ws://localhost/8080/ws/winds

建立连接后发送json数据做为请求

格式

```json
{"devices": ["Q"], "limit": 50, "enable": true, "close": true}
```

- devices表示需要的设备

- enable: 为True激活websocket会发数据,False禁用Websocket回传数据

- close: 为True时关闭websocket，也可以直接关闭websocket连接

返回数据格式

历史数据与/winds api一致

单条实时数据

```json
{"Id":2550,"DeviceId":"Q","Direction":229,"Speed":2.74,"Unit":"M","CreateAt":"2020-04-25T16:19:34.7598258Z"}
```