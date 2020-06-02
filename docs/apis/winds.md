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

### 最大速度接口

```
http://101.132.38.140:8080/api/speeds?period=170&interval=240
```

```json
{
  "A": [  // A为设备ID
    {
      "Speed": 0.08,
      "CreateAt": "2020-05-26T23:01:03+08:00"
    },
    {
      "Speed": 0.09,
      "CreateAt": "2020-05-27T02:10:00+08:00"
    },
    {
      "Speed": 0.23,
      "CreateAt": "2020-05-27T08:28:26+08:00"
    },
    {
      "Speed": 0.27,
      "CreateAt": "2020-05-27T10:15:26+08:00"
    },
    {
      "Speed": 0.24,
      "CreateAt": "2020-05-27T17:46:36+08:00"
    }
  ],
  "Q": [
    {
      "Speed": 0.12,
      "CreateAt": "2020-05-27T11:10:59+08:00"
    },
    {
      "Speed": 0.1,
      "CreateAt": "2020-05-27T14:27:25+08:00"
    },
    {
      "Speed": 0.06,
      "CreateAt": "2020-05-27T22:01:38+08:00"
    },
    {
      "Speed": 0.09,
      "CreateAt": "2020-05-28T01:51:12+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-28T02:41:55+08:00"
    },
    {
      "Speed": 0.14,
      "CreateAt": "2020-05-28T09:07:51+08:00"
    },
    {
      "Speed": 0.11,
      "CreateAt": "2020-05-28T10:25:51+08:00"
    },
    {
      "Speed": 0.11,
      "CreateAt": "2020-05-28T14:12:48+08:00"
    },
    {
      "Speed": 0.05,
      "CreateAt": "2020-05-28T18:09:10+08:00"
    },
    {
      "Speed": 0.05,
      "CreateAt": "2020-05-28T22:31:19+08:00"
    },
    {
      "Speed": 0.06,
      "CreateAt": "2020-05-29T05:26:19+08:00"
    },
    {
      "Speed": 0.05,
      "CreateAt": "2020-05-29T06:16:42+08:00"
    },
    {
      "Speed": 0.06,
      "CreateAt": "2020-05-29T12:19:49+08:00"
    },
    {
      "Speed": 0.05,
      "CreateAt": "2020-05-29T14:56:31+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-29T18:52:40+08:00"
    },
    {
      "Speed": 0.08,
      "CreateAt": "2020-05-30T00:48:43+08:00"
    },
    {
      "Speed": 0.08,
      "CreateAt": "2020-05-30T05:45:56+08:00"
    },
    {
      "Speed": 0.08,
      "CreateAt": "2020-05-30T06:59:25+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-30T10:17:04+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-30T14:34:24+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-30T19:00:41+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-30T22:17:35+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-31T02:40:58+08:00"
    },
    {
      "Speed": 0.07,
      "CreateAt": "2020-05-31T09:55:24+08:00"
    },
    {
      "Speed": 0.08,
      "CreateAt": "2020-05-31T12:40:12+08:00"
    },
    {
      "Speed": 0.18,
      "CreateAt": "2020-05-31T16:33:51+08:00"
    },
    {
      "Speed": 0.76,
      "CreateAt": "2020-05-31T19:58:33+08:00"
    }
  ]
}
```
