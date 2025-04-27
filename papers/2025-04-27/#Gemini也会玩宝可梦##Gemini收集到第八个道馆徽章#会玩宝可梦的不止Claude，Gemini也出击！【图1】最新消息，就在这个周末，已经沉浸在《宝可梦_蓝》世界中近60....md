# #Gemini也会玩宝可梦##Gemini收集到第八个道馆徽章#会玩宝可梦的不止Claude，Gemini也出击！【图1】最新消息，就在这个周末，已经沉浸在《宝可梦 蓝》世界中近60...

**URL**: https://weibo.com/6105753431/PpdpF3PGv

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Gemini%E4%B9%9F%E4%BC%9A%E7%8E%A9%E5%AE%9D%E5%8F%AF%E6%A2%A6%23&amp;extparam=%23Gemini%E4%B9%9F%E4%BC%9A%E7%8E%A9%E5%AE%9D%E5%8F%AF%E6%A2%A6%23" data-hide=""><span class="surl-text">#Gemini也会玩宝可梦#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Gemini%E6%94%B6%E9%9B%86%E5%88%B0%E7%AC%AC%E5%85%AB%E4%B8%AA%E9%81%93%E9%A6%86%E5%BE%BD%E7%AB%A0%23&amp;extparam=%23Gemini%E6%94%B6%E9%9B%86%E5%88%B0%E7%AC%AC%E5%85%AB%E4%B8%AA%E9%81%93%E9%A6%86%E5%BE%BD%E7%AB%A0%23" data-hide=""><span class="surl-text">#Gemini收集到第八个道馆徽章#</span></a><br><br>会玩宝可梦的不止Claude，Gemini也出击！【图1】<br><br>最新消息，就在这个周末，已经沉浸在《宝可梦 蓝》世界中近600个小时的Gemini 2.5 Pro成功收集到了第八个道馆徽章！【图2】<br><br>接下来，它将开启冠军之路&amp;四天王之战！<br><br>得知了这个消息的网友们纷纷传来贺电：“这是AI迄今为止取得的最大成就！”【图3】<br><br>这次挑战的起因，来自先前的一个类似项目：让Claude挑战宝可梦。【图4】<br><br>一名30岁的软件工程师在观看了这个项目后灵机一动，让Gemini也踏上了成为最强的宝可梦训练家之路。<br><br>为了让Gemini顺利地进行游戏，这名工程师介绍了他的核心思路：<br><br>一、核心组件<br>游戏接口：<br>代码通过套接字连接与运行《宝可梦蓝》的mGBA模拟器建立通信，使其能够：<br>- 截取当前游戏状态的屏幕截图<br>- 发送按键指令<br>- 获取游戏状态数据，包括玩家位置、宝可梦队伍信息和地图详情<br><br>决策系统：<br>该系统会对游戏截图进行网格叠加处理，并将这些截图（连连同从游戏内存中提取的状态信息）发送给AI进行分析，由AI根据获得的信息决定下一步操作。<br><br>二、游戏循环<br>在游戏中，以下操作将会被循环进行：<br>1. 捕获屏幕截图并获取游戏状态数据<br>2. 对图像进行网格叠加处理以辅助空间推理<br>3. 将截图和游戏信息发送至AI模型<br>4. 解析AI响应以确定需要按下的按键<br>5. 执行按键操作并等待游戏状态更新<br>6. 为下一帧画面重复此流程<br><br>三、内存管理<br>当前系统约每执行100次动作后会对消息进行摘要处理，以此减少输入token的消耗。摘要内容将替换原始消息。<br><br>目前，Gemini的挑战还在继续，大家可以进入直播间围观一波。<br><br>童年时的你估计也想不到，有朝一日居然会在互联网上收看AI玩宝可梦吧！<br><br>直播链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.twitch.tv%2Fgemini_plays_pokemon" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0vi4mppadj33gc1xyu0y.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0vi4myy6ij30e80ap3zw.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0vi4qrjmgj30ww16aame.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0vi4z3eymj33gg1y0npf.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Gemini AI成功在《宝可梦 蓝》中收集到第八个道馆徽章，耗时近600小时。一名工程师通过mGBA模拟器建立游戏接口，让Gemini 2.5 Pro能获取游戏截图和状态数据，经网格处理和AI分析后执行按键操作。系统每100次动作会摘要消息以减少token消耗。该项目灵感来自Claude玩宝可梦的尝试，现已在Twitch直播AI挑战冠军之路，引发网友热议。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-27T13:08:02Z
- **目录日期**: 2025-04-27
