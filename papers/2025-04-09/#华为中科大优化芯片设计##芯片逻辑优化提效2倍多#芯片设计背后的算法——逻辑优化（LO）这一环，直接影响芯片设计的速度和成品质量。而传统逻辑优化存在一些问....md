# #华为中科大优化芯片设计##芯片逻辑优化提效2倍多#芯片设计背后的算法——逻辑优化（LO）这一环，直接影响芯片设计的速度和成品质量。而传统逻辑优化存在一些问...

**URL**: https://weibo.com/6105753431/PmtGny0ie

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E4%B8%AD%E7%A7%91%E5%A4%A7%E4%BC%98%E5%8C%96%E8%8A%AF%E7%89%87%E8%AE%BE%E8%AE%A1%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E4%B8%AD%E7%A7%91%E5%A4%A7%E4%BC%98%E5%8C%96%E8%8A%AF%E7%89%87%E8%AE%BE%E8%AE%A1%23" data-hide=""><span class="surl-text">#华为中科大优化芯片设计#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8A%AF%E7%89%87%E9%80%BB%E8%BE%91%E4%BC%98%E5%8C%96%E6%8F%90%E6%95%882%E5%80%8D%E5%A4%9A%23&amp;extparam=%23%E8%8A%AF%E7%89%87%E9%80%BB%E8%BE%91%E4%BC%98%E5%8C%96%E6%8F%90%E6%95%882%E5%80%8D%E5%A4%9A%23" data-hide=""><span class="surl-text">#芯片逻辑优化提效2倍多#</span></a><br><br>芯片设计背后的算法——逻辑优化（LO）这一环，直接影响芯片设计的速度和成品质量。<br><br>而传统逻辑优化存在一些问题：冗余操作太多，跑得慢，还经常“白忙活”。<br><br>为了优化这个环节，中科大王杰教授团队和华为诺亚方舟实验室联手，搞出了一个新AI框架——CMO（Circuit Symbolic Learning Framework）。<br><br>它结合图神经网络（GNN）和蒙特卡洛树搜索（MCTS），能大幅减少无效操作，“剪枝”更精准，让芯片设计效率最高提升了2.5倍。<br><br>打个比方，GNN像是“看图识路”的AI司机，能理解电路图的结构，而MCTS就像是军师，帮你避开坑点、挑出高收益操作。<br><br>CMO把这两者结合在一起，不光跑得快，还跑得准，而且能用在只靠CPU的实际工业环境里，告别GPU依赖。<br><br>更妙的是，这不是简单粗暴地堆深度学习，而是搞了个“教师-学生”模式：GNN先学明白逻辑转换哪些有效，再把这个“知识”用符号学习的方法传授给MCTS，最后产出轻量级、可解释、可泛化的打分函数。<br><br>CMO中最核心的一招叫GESD（Graph Enhanced Symbolic Discovery Framework），这是一种图增强的蒙特卡洛树搜索方案。<br><br>它定义了一套规则和符号（比如加减乘除、对数、布尔运算等），再通过蒙特卡洛树一步步生成“表达式树”，不断尝试不同的组合，找出最能判断节点转换价值的函数。<br><br>它还用了多种策略优化整个流程：<br><br>- 结构-语义特征分离：结构特征走数学路线，语义特征用布尔逻辑，各走各的，避免搜索空间爆炸；<br>- 知识蒸馏：不是照搬GNN的结果，而是学习GNN的“判断思路”；<br>- 奖励机制调优：确保学习出的函数既有效又不复杂，兼顾实用性和效率；<br>- 焦点损失函数（Focal Loss）：针对稀疏数据中正负样本不平衡问题，让训练更稳准狠。<br><br>实验结果来看，在EPFL和IWLS两个开源电路数据集，以及一个工业级电路数据集上验证，CMO不仅让老牌逻辑优化算子Mfs2提速，还能在相同时间内跑出更好的电路优化结果。<br><br>像在大规模Sixteen电路上，优化时间从78784秒缩短到32001秒，效率提升接近60%。<br><br>更惊喜的是，电路深度也减了不少，比如在Hyp电路上从8259层减少到5762层，大幅降低延迟，真正让“剪枝”有感。<br><br>目前，这套技术已经集成进华为的EMU逻辑综合工具中，成为支持EDA国产化链条的重要一环。<br><br>原文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2Fwf8XgLrVpov4ll7Ha7ef-Q" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>论文地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fopenreview.net%2Fforum%3Fid%3DEG9nDN3eGB" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>代码地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgitee.com%2Fyinqi-bai%2Fcmo" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0aou7almuj31dm114dvq.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

中科大与华为合作开发了新型AI框架CMO，用于优化芯片设计中的逻辑优化（LO）环节。该框架结合图神经网络（GNN）和蒙特卡洛树搜索（MCTS），通过"教师-学生"模式生成轻量级、可解释的评分函数，显著减少无效操作。实验显示，CMO将优化效率提升最高2.5倍，如Sixteen电路优化时间从78784秒缩短至32001秒，电路深度也大幅降低。该技术已集成到华为EMU工具中，支持EDA国产化。研究采用GESD方案，通过结构-语义特征分离、知识蒸馏等方法优化流程。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T21:03:26Z
- **目录日期**: 2025-04-09
