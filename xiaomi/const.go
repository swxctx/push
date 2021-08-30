package xiaomi

const (
	// 基础请求url
	BaseReqUrl = "https://api.xmpush.xiaomi.com"
)

const (
	// 表示通知栏消息
	PASS_THROUGH_BY_NOTICE = 0
	// 表示透传消息
	PASS_THROUGH_BY_TCP = 1

	// 所有类型
	NOTIFY_TYPE_DEFAULT_ALL = -1
	// 使用默认提示音提示
	NOTIFY_TYPE_DEFAULT_SOUND = 1
	// 使用默认震动提示
	NOTIFY_TYPE_DEFAULT_VIBRATE = 2
	// 使用默认led灯光提示
	NOTIFY_TYPE_DEFAULT_LIGHTS = 4

	// 允许弹出通知栏消息
	NOTIFY_FOREGROUND_YES = 1
	// 不允许弹出通知栏消息
	NOTIFY_FOREGROUND_NO = 0
)
