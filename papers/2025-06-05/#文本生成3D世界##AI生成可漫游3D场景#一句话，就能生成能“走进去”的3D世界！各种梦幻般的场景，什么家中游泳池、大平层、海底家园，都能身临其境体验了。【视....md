# #文本生成3D世界##AI生成可漫游3D场景#一句话，就能生成能“走进去”的3D世界！各种梦幻般的场景，什么家中游泳池、大平层、海底家园，都能身临其境体验了。【视...

**URL**: https://weibo.com/6105753431/Pv8aa6Z0B

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%96%87%E6%9C%AC%E7%94%9F%E6%88%903D%E4%B8%96%E7%95%8C%23&amp;extparam=%23%E6%96%87%E6%9C%AC%E7%94%9F%E6%88%903D%E4%B8%96%E7%95%8C%23" data-hide=""><span class="surl-text">#文本生成3D世界#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E7%94%9F%E6%88%90%E5%8F%AF%E6%BC%AB%E6%B8%B83D%E5%9C%BA%E6%99%AF%23&amp;extparam=%23AI%E7%94%9F%E6%88%90%E5%8F%AF%E6%BC%AB%E6%B8%B83D%E5%9C%BA%E6%99%AF%23" data-hide=""><span class="surl-text">#AI生成可漫游3D场景#</span></a><br><br>一句话，就能生成能“走进去”的3D世界！<br><br>各种梦幻般的场景，什么家中游泳池、大平层、海底家园，都能身临其境体验了。【视频】<br><br>这就是慕尼黑工业大学提出的WorldExplorer，不同于以往只能“站在原地看”的3D生成方法，WorldExplorer实现了真正的空间探索能力，背后有三大关键机制：<br><br>WorldExplorer技术上的几大步骤包括：<br><br>- 全景初始化：用文本生成4张外向图像，通过单目深度估计和图像修复技术，再生成4张补全视图，构建一个完整360°场景框架。<br>    <br>- 迭代式视频拓展：用摄像头引导的视频扩散模型，沿预设路径从不同初始图出发，生成一批短视频片段，探索场景内部和物体周围。每次生成都依据历史图像选取最相关的视角进行条件控制。<br>    <br>- 3D建模优化：将所有图像输入Gaussian Splatting优化器，先通过VGGT算法构建稀疏点云初始化，再进行相机对齐，最终生成可实时渲染、可自由漫游的完整3D场景。<br><br>进一步来说，每次视频生成都会参考场景记忆库中，与当前路径旋转角度最接近的5帧，外加最初的8帧全景框架。这样能确保新生成图像与历史保持空间一致性，同时提升结构连贯性。<br><br>为避免生成“穿墙”或“贴脸”画面，系统还会提前检测相机路径上是否会“撞上”物体，一旦检测到潜在碰撞，就自动丢弃该帧，从源头避免畸变与破图。<br><br>WorldExplorer的最终输出可用于实时渲染，不仅能还原复杂场景细节，还支持大范围自由移动，解决了传统文本生成3D方法中“只能在中间看”的问题。<br><br>项目主页：the-world-explorer.github.io/<img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i24h8d56u8j30k00k0q3i.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i24h8ext3bj30u00u0mys.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i24h8hph5dj30u00u0gms.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i24h6c5zkaj30zk0q2tu3.jpg" referrerpolicy="no-referrer"><br><br><br clear="both"><div style="clear: both"></div><video controls="controls" poster="https://tvax2.sinaimg.cn/orj480/006Fd7o3ly1i24h8d7duoj30k00k0q3i.jpg" style="width: 100%"><source src="https://f.video.weibocdn.com/o0/RmkERjaJlx08oNSaDvmg01041200b9wr0E010.mp4?label=mp4_720p&amp;template=720x720.24.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1749113903&amp;ssig=WOQJ0AQOXq&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/8mYI63ahlx08oNSaMpao010412007Ca80E010.mp4?label=mp4_hd&amp;template=540x540.24.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1749113903&amp;ssig=9K4AiklFG7&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/m4q63Q0Zlx08oNSa6Cko010412004uLs0E010.mp4?label=mp4_ld&amp;template=360x360.24.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1749113903&amp;ssig=TbG9%2Fi%2F%2BzI&amp;KID=unistore,video"><p>视频无法显示，请前往<a href="https://video.weibo.com/show?fid=1034%3A5174193957044324" target="_blank" rel="noopener noreferrer">微博视频</a>观看。</p></video>

## AI 摘要

慕尼黑工业大学开发的WorldExplorer系统实现了通过文本生成可自由漫游的3D场景。该系统通过三个关键步骤：1) 用文本生成4张外向图像并补全360°全景框架；2) 使用视频扩散模型沿路径生成探索性短视频片段；3) 采用Gaussian Splatting技术优化3D建模，最终输出可实时渲染的完整3D场景。创新性地解决了传统方法只能固定视角的问题，通过场景记忆库保持空间一致性，并自动检测避免穿墙等异常情况。用户仅需输入文字描述，即可生成如家中泳池、海底世界等可交互探索的3D环境。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T08:04:20Z
- **目录日期**: 2025-06-05
