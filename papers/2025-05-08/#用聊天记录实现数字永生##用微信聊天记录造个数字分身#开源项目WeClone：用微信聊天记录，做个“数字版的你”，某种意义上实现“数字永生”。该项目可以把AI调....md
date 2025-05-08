# #用聊天记录实现数字永生##用微信聊天记录造个数字分身#开源项目WeClone：用微信聊天记录，做个“数字版的你”，某种意义上实现“数字永生”。该项目可以把AI调...

**URL**: https://weibo.com/6105753431/PqJEQur0w

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8%E8%81%8A%E5%A4%A9%E8%AE%B0%E5%BD%95%E5%AE%9E%E7%8E%B0%E6%95%B0%E5%AD%97%E6%B0%B8%E7%94%9F%23&amp;extparam=%23%E7%94%A8%E8%81%8A%E5%A4%A9%E8%AE%B0%E5%BD%95%E5%AE%9E%E7%8E%B0%E6%95%B0%E5%AD%97%E6%B0%B8%E7%94%9F%23" data-hide=""><span class="surl-text">#用聊天记录实现数字永生#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8%E5%BE%AE%E4%BF%A1%E8%81%8A%E5%A4%A9%E8%AE%B0%E5%BD%95%E9%80%A0%E4%B8%AA%E6%95%B0%E5%AD%97%E5%88%86%E8%BA%AB%23&amp;extparam=%23%E7%94%A8%E5%BE%AE%E4%BF%A1%E8%81%8A%E5%A4%A9%E8%AE%B0%E5%BD%95%E9%80%A0%E4%B8%AA%E6%95%B0%E5%AD%97%E5%88%86%E8%BA%AB%23" data-hide=""><span class="surl-text">#用微信聊天记录造个数字分身#</span></a><br><br>开源项目WeClone：用微信聊天记录，做个“数字版的你”，某种意义上实现“数字永生”。<br><br>该项目可以把AI调成模仿你的口气说话，再接入聊天机器人，实现属于你自己的“数字分身”。<br><br>来看下它具体能干嘛：<br><br>- 支持从微信导出聊天记录，并自动处理成问答格式<br>- 基于LoRA方法、微调Qwen2.5-7B等模型，让LLM说话更像你<br>- 还能克隆语音，基于0.5B模型重现你的语气语调（配套子项目：WeClone-audio）<br>- 最终结果可以部署到微信/QQ/飞书/企微等多平台，实现聊天式交互<br><br>WeClone也贴心地给出了详细的流程文档，从环境配置、数据处理 到 模型训练、服务部署都有说明。<br><br>还支持单卡/多卡训练、FlashAttention加速、LoRA参数配置等技术细节，适合技术向用户上手。<br><br>不过要注意几点：<br><br>- 微调效果依赖聊天数据的质量和模型大小<br>- Windows下建议用WSL（Linux内核）<br>- 当前版本正在快速迭代中，不稳定属正常现象<br>- 涉及隐私数据，一定要合法合规使用<br><br>感兴趣的可以上GitHub看源码：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fxming521%2FWeClone" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a> <a href="https://video.weibo.com/show?fid=1034:5163720041758750" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_video_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">量子位的微博视频</span></a><br clear="both"><div style="clear: both"></div><video controls="controls" poster="https://tvax4.sinaimg.cn/orj480/006Fd7o3ly1i1727g8bt6j31690u0q52.jpg" style="width: 100%"><source src="https://f.video.weibocdn.com/o0/U8HYgkyClx08o3LGPs1i010412003YMl0E010.mp4?label=mp4_720p&amp;template=1012x720.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1746666138&amp;ssig=I5Ed8yNRj7&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/WRZTdDyUlx08o3LGZOCs010412001SSW0E010.mp4?label=mp4_hd&amp;template=676x480.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1746666138&amp;ssig=sC2VfHw%2B7q&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/93IwzWbClx08o3LH1KhO0104120018aX0E010.mp4?label=mp4_ld&amp;template=504x360.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1746666138&amp;ssig=5fvvhc5bj0&amp;KID=unistore,video"><p>视频无法显示，请前往<a href="https://video.weibo.com/show?fid=1034%3A5163720041758750" target="_blank" rel="noopener noreferrer">微博视频</a>观看。</p></video>

## AI 摘要

WeClone是一个开源项目，通过分析微信聊天记录训练AI模型，生成用户的"数字分身"。该项目支持导出聊天记录并转为问答格式，基于LoRA方法微调Qwen2.5-7B等大模型，使AI模仿用户说话风格，还能克隆语音语调。最终可将数字分身部署到微信/QQ等平台。项目提供详细技术文档，支持单卡/多卡训练和FlashAttention加速，但效果取决于数据质量和模型大小，需注意隐私合规问题。目前仍在快速迭代中，适合技术用户尝试。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T00:03:17Z
- **目录日期**: 2025-05-08
