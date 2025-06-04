# #模型顿悟原来是因为记不住了##模型记忆总量有固定上限#你有没好奇过，语言模型究竟能记住多少信息？Meta的最新研究带来了答案：GPT类模型，每个参数能记忆3.6比...

**URL**: https://weibo.com/6105753431/PuYuoaDer

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%A8%A1%E5%9E%8B%E9%A1%BF%E6%82%9F%E5%8E%9F%E6%9D%A5%E6%98%AF%E5%9B%A0%E4%B8%BA%E8%AE%B0%E4%B8%8D%E4%BD%8F%E4%BA%86%23&amp;extparam=%23%E6%A8%A1%E5%9E%8B%E9%A1%BF%E6%82%9F%E5%8E%9F%E6%9D%A5%E6%98%AF%E5%9B%A0%E4%B8%BA%E8%AE%B0%E4%B8%8D%E4%BD%8F%E4%BA%86%23" data-hide=""><span class="surl-text">#模型顿悟原来是因为记不住了#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%A8%A1%E5%9E%8B%E8%AE%B0%E5%BF%86%E6%80%BB%E9%87%8F%E6%9C%89%E5%9B%BA%E5%AE%9A%E4%B8%8A%E9%99%90%23&amp;extparam=%23%E6%A8%A1%E5%9E%8B%E8%AE%B0%E5%BF%86%E6%80%BB%E9%87%8F%E6%9C%89%E5%9B%BA%E5%AE%9A%E4%B8%8A%E9%99%90%23" data-hide=""><span class="surl-text">#模型记忆总量有固定上限#</span></a><br><br>你有没好奇过，语言模型究竟能记住多少信息？<br><br>Meta的最新研究带来了答案：GPT类模型，每个参数能记忆3.6比特数据。【图1】<br><br>并且，在通过随机均匀字符串进行约100万步的超大批量训练，将所有模型训练至“饱和”状态之后，研究团队还发现个很有趣的现象：<br><br>模型记忆总量与训练数据规模无关。这意味着模型本身有一个固定的“记忆上限”，即使给它再多的样本，它的记忆也会“摊薄”，无法记住所有细节。【图2】<br><br>这个发现也很好地解释了模型为什么会突然“顿悟”：当模型的记忆容量饱和，无法完美记住所有训练样本时，它就必须开始智能地共享样本之间的信息。【图3】【图4】<br><br>当模型用文本数据进行训练时，情况会稍有不同：它只会记住参数容量允许范围内的样本。一旦超过这个记忆极限，模型就会放弃逐个样本的记忆，转而共享信息，这就是我们常说的“泛化”。【图5】<br><br>研究团队在完全去重、纯净的文本数据环境中进行了实验，还发现了一些关于隐私的重要结论：<br><br>- 当模型记忆容量充分饱和后，测试样本反而比训练样本更容易被提取出来。【图6】<br>- 最容易被提取的是那些包含极其罕见token的样本<br>- 成员推理远比数据提取容易【图7】<br><br>最终，研究团队得出结论：接受海量数据训练的模型无法记忆全部训练数据。根本原因在于容量不足。【图8】<br><br>Meta团队的研究人员究竟是如何准确计算记忆容量数据的呢？这背后具有两项关键创新：<br><br>1. 排除泛化的影响<br><br>研究团体将记忆现象划分为“非预期记忆”和“泛化”，泛化是模型对真实数据生成规律的理解。<br><br>研究团队用没有内在规律的随机均匀字符串进行测试。这样就迫使模型只能通过记忆数据点来降低错误。<br><br>2. 用“压缩率”量化模型记忆程度<br><br>团队使用Kolmogorov复杂度理论来定义和测量模型对数据的记忆，计算出了模型在给定数据集上的压缩率。<br><br>你可以这样理解：模型对数据的记忆量，就是原始数据的“混乱程度”（熵）减去模型压缩后数据的“混乱程度”。<br><br>在随机均匀字符串数据集上，由于数据完全随机，其“混乱程度”是已知的，所以可以直接计算出模型对数据的记忆量。<img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i23aj40ngcj30mg0m2gph.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i23aj5woe7j30mg0hc43e.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i23aj8nsxjj30n80n4tdc.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i23aj9lwfyj30n40lugri.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i23ajbr9q6j30n40n60y9.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i23ajfc3u6j30iu0jyq6l.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i23ajh1zezj30iq0kwq61.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i23ajjdcpgj30r60qajz0.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Meta研究发现，GPT类模型每个参数仅能记忆3.6比特数据，且记忆总量存在固定上限。当模型记忆饱和后，会通过信息共享实现"泛化"，这解释了其"顿悟"现象。研究采用随机字符串训练排除泛化干扰，并用Kolmogorov复杂度量化记忆量，发现：1）测试样本比训练样本更易被提取；2）罕见token样本最易泄露；3）模型因容量限制无法记住全部训练数据。该研究揭示了语言模型的记忆机制及其隐私风险。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T07:01:21Z
- **目录日期**: 2025-06-04
