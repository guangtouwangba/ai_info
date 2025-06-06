# #智能体本质竟是世界模型##DeepMind证明Ilya神预言#DeepMind科学家Jon Richens提出一个有趣观点：智能体本身，就是世界模型。【图1】他甩出一篇ICML 2025论文，...

**URL**: https://weibo.com/6105753431/PvimsEohx

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%99%BA%E8%83%BD%E4%BD%93%E6%9C%AC%E8%B4%A8%E7%AB%9F%E6%98%AF%E4%B8%96%E7%95%8C%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E6%99%BA%E8%83%BD%E4%BD%93%E6%9C%AC%E8%B4%A8%E7%AB%9F%E6%98%AF%E4%B8%96%E7%95%8C%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#智能体本质竟是世界模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23DeepMind%E8%AF%81%E6%98%8EIlya%E7%A5%9E%E9%A2%84%E8%A8%80%23&amp;extparam=%23DeepMind%E8%AF%81%E6%98%8EIlya%E7%A5%9E%E9%A2%84%E8%A8%80%23" data-hide=""><span class="surl-text">#DeepMind证明Ilya神预言#</span></a><br><br>DeepMind科学家Jon Richens提出一个有趣观点：智能体本身，就是世界模型。【图1】<br><br>他甩出一篇ICML 2025论文，文章意思是，只要一个智能体能泛化完成复杂任务，那它策略里一定已经学出了环境的内部结构。<br><br>这个结论，神同步了OpenAI的Ilya在2023年说的一句话：“世界模型，可能是所有智能体背后那条隐藏的自然法则。”【视频2】<br><br>而且另一篇RLC 2025文章（投稿中）提出的观点，也侧面印证了这种说法。【图3】<br><br>回到这篇论文中来，DeepMind团队从“后悔值界限”出发（后悔值越小，智能体越接近最优），证明：<br><br>- 如果一个智能体在很多目标导向任务上表现不差；<br><br>- 并且这些任务跨越多个时间步、需要一定推理；<br><br>- 那它的行为策略中，必然隐含了一个可恢复的环境模型（即世界模型）；<br><br>- 换句话说：想让策略好用，就必须“了解世界”。<br><br>特别地，他们还提出了一种新算法，可以直接从“目标+智能体策略”中，反推出环境的转移规律（P(s'|s,a)）。<br><br>这就补上了智能学习三角的最后一角：<br><br>- 规划：世界模型+目标→策略<br>- 逆强化学习：世界模型+策略→目标<br>- 本研究：策略+目标→世界模型<br><br>这个发现，直接打破了许多人的“无模型”幻想——<br><br>很多人希望靠“大力出奇迹”的方式直接训练智能体，不建世界模型、只优化回报。但这项研究证明，只要任务足够复杂，这种“无模型捷径”根本走不通。<br><br>哪怕是今天的LLM或Gato、PaLM-E这类多任务智能体，它们的成功，其实是因为在背后学出了某种世界预测机制。<br><br>即，只要目标设置得好，就能诱导智能体不断“推演”世界。这个观点也给智能体训练方式带来新灵感：不必靠环境建模，而是用目标反向逼出世界规律。<br><br>这项研究不是告诉我们“AGI需要世界模型”这么简单，而是：没有世界模型，就不可能有具备泛化能力的智能体。<br><br>这不仅是一种必要条件，也是一种更底层的认知框架。<br><br>它意味着，我们未来构建AI系统，不一定非得手动建模世界——但一定得通过目标驱动或策略涌现的方式，让AI“自动学出”世界模型。<br><br>也许，这就是“智能体”的本质。<img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i25q80mte5j314w11kndg.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i25q9kfow7j30zk0k0jsp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i25q8o2qdyj310218wtux.jpg" referrerpolicy="no-referrer"><br><br><br clear="both"><div style="clear: both"></div><video controls="controls" poster="https://tvax2.sinaimg.cn/orj480/006Fd7o3ly1i25q9kii13j30zk0k0jsp.jpg" style="width: 100%"><source src="https://f.video.weibocdn.com/o0/4Y78WGtFlx08oPBb3CbS01041200pzhw0E010.mp4?label=mp4_720p&amp;template=1280x720.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1749214966&amp;ssig=m3Os8itNG7&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/bF8hmOW5lx08oPBaFO9G01041200cL8S0E010.mp4?label=mp4_hd&amp;template=852x480.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1749214966&amp;ssig=0yCVD41yIr&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/BtdeUwJUlx08oPBaGYdq010412008mj20E010.mp4?label=mp4_ld&amp;template=640x360.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1749214966&amp;ssig=s4I7D%2Fp1W9&amp;KID=unistore,video"><p>视频无法显示，请前往<a href="https://video.weibo.com/show?fid=1034%3A5174587151810576" target="_blank" rel="noopener noreferrer">微博视频</a>观看。</p></video>

## AI 摘要

DeepMind最新研究证实，智能体的本质就是世界模型。研究发现，能泛化处理复杂任务的智能体，其策略必然隐含对环境内部结构的建模。该结论印证了OpenAI首席科学家Ilya Sutskever的预言，并填补了"策略+目标→世界模型"的智能学习三角闭环。研究打破了"无模型强化学习"的幻想，表明即使是当前大语言模型，其成功也源于隐含的世界预测机制。这为AI开发提供了新思路：通过目标驱动让系统自动构建世界模型，而非人工建模。该发现确立了世界模型作为智能体泛化能力的必要条件。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T12:03:06Z
- **目录日期**: 2025-06-06
