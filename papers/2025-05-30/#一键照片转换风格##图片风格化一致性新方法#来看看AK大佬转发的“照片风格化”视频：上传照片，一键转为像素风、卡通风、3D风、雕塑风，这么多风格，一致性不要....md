# #一键照片转换风格##图片风格化一致性新方法#来看看AK大佬转发的“照片风格化”视频：上传照片，一键转为像素风、卡通风、3D风、雕塑风，这么多风格，一致性不要...

**URL**: https://weibo.com/6105753431/PudfFAq3v

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E9%94%AE%E7%85%A7%E7%89%87%E8%BD%AC%E6%8D%A2%E9%A3%8E%E6%A0%BC%23&amp;extparam=%23%E4%B8%80%E9%94%AE%E7%85%A7%E7%89%87%E8%BD%AC%E6%8D%A2%E9%A3%8E%E6%A0%BC%23" data-hide=""><span class="surl-text">#一键照片转换风格#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%9B%BE%E7%89%87%E9%A3%8E%E6%A0%BC%E5%8C%96%E4%B8%80%E8%87%B4%E6%80%A7%E6%96%B0%E6%96%B9%E6%B3%95%23&amp;extparam=%23%E5%9B%BE%E7%89%87%E9%A3%8E%E6%A0%BC%E5%8C%96%E4%B8%80%E8%87%B4%E6%80%A7%E6%96%B0%E6%96%B9%E6%B3%95%23" data-hide=""><span class="surl-text">#图片风格化一致性新方法#</span></a><br><br>来看看AK大佬转发的“照片风格化”视频：上传照片，一键转为像素风、卡通风、3D风、雕塑风，这么多风格，一致性不要太稳。（文末有体验网址）<br><br>要知道，风格化一直有两个问题：一个是风格不稳，另一个是结构容易乱。<br><br>而这个视频，用到了OmniConsistency方法——它是一个“可插拔”的一致性模块，可以搭配各种扩散式风格迁移模型使用，只加个轻模块，就能显著提升效果。<br><br>背后的逻辑包括：<br><br>- 两阶段训练：先单独训练风格LoRA，再用风格前后的图像对来训练一致性模块，避免风格干扰结构。<br><br>- 滚动LoRA策略：训练时每隔50步换一个风格模块和数据集，让模型能适应不同风格，包括没见过的新风格。<br><br>- 一致性LoRA：只加在条件分支上，不动主干，既保持结构信息，也不影响风格表达。<br><br>- 位置感知插值+因果注意力：保证低分辨率条件图能正确指导高分辨率输出，同时不乱串信息。<br><br>- 数据集也不马虎：用GPT-4o生成的风格图像+人工挑选，22种风格、2600对高质量数据。<br><br>结果上也很能打。无论是DreamSim、FID、还是GPT-4o评分，各项指标都对标甚至超过了现有SOTA，用户偏好测试也更倾向于OmniConsistency在风格和内容一致性上的表现。<br><br>虽然还有些边角问题，比如非英文文本处理不够好、多人物场景下细节不稳，但整体来说，它已经是目前风格迁移里“结构不崩、风格不丢”的强解法之一。<br><br>而且对接起来也方便，只加5%左右的资源开销，就能适配各种现有LoRA和风格框架（如IP-Adapter）。<br><br>在线网站：huggingface.co/spaces/yiren98/OmniConsistency<br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.18445" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1xi01umezj30k00k0gmt.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1xhz4vk3wj30uo0rue3i.jpg" referrerpolicy="no-referrer"><br><br><br clear="both"><div style="clear: both"></div><video controls="controls" poster="https://tvax3.sinaimg.cn/orj480/006Fd7o3ly1i1xi01yb5nj30k00k0gmt.jpg" style="width: 100%"><source src="https://f.video.weibocdn.com/o0/rydpuqXjlx08oEfdseeY010412004ecg0E010.mp4?label=mp4_720p&amp;template=720x720.24.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1748595756&amp;ssig=HU8oQHQNNt&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/dbZGscXvlx08oEfd8mtW010412002OAE0E010.mp4?label=mp4_hd&amp;template=540x540.24.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1748595756&amp;ssig=Hxhe%2BDoIN8&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/sIMaP9rolx08oEfcSe24010412001BHA0E010.mp4?label=mp4_ld&amp;template=360x360.24.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1748595756&amp;ssig=dOm4ykwjJS&amp;KID=unistore,video"><p>视频无法显示，请前往<a href="https://video.weibo.com/show?fid=1034%3A5172007533215809" target="_blank" rel="noopener noreferrer">微博视频</a>观看。</p></video>

## AI 摘要

OmniConsistency是一种创新的图片风格化一致性方法，通过可插拔模块显著提升扩散模型的效果。其核心技术包括两阶段训练、滚动LoRA策略和位置感知插值，确保结构稳定且风格多样。使用GPT-4o生成的高质量数据集（22种风格/2600对数据），在DreamSim、FID等指标上超越现有SOTA，用户偏好测试也验证了其优越性。尽管存在非英文文本处理等小问题，但仅增加5%资源开销即可适配主流框架。体验地址：Hugging Face空间；论文见arXiv链接。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T08:04:16Z
- **目录日期**: 2025-05-30
