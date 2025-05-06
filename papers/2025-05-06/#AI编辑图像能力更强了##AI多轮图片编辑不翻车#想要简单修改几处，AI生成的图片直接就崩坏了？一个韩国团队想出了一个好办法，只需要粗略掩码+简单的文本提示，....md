# #AI编辑图像能力更强了##AI多轮图片编辑不翻车#想要简单修改几处，AI生成的图片直接就崩坏了？一个韩国团队想出了一个好办法，只需要粗略掩码+简单的文本提示，...

**URL**: https://weibo.com/6105753431/PqxIJgivT

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E7%BC%96%E8%BE%91%E5%9B%BE%E5%83%8F%E8%83%BD%E5%8A%9B%E6%9B%B4%E5%BC%BA%E4%BA%86%23&amp;extparam=%23AI%E7%BC%96%E8%BE%91%E5%9B%BE%E5%83%8F%E8%83%BD%E5%8A%9B%E6%9B%B4%E5%BC%BA%E4%BA%86%23" data-hide=""><span class="surl-text">#AI编辑图像能力更强了#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%A4%9A%E8%BD%AE%E5%9B%BE%E7%89%87%E7%BC%96%E8%BE%91%E4%B8%8D%E7%BF%BB%E8%BD%A6%23&amp;extparam=%23AI%E5%A4%9A%E8%BD%AE%E5%9B%BE%E7%89%87%E7%BC%96%E8%BE%91%E4%B8%8D%E7%BF%BB%E8%BD%A6%23" data-hide=""><span class="surl-text">#AI多轮图片编辑不翻车#</span></a><br><br>想要简单修改几处，AI生成的图片直接就崩坏了？<br><br>一个韩国团队想出了一个好办法，只需要粗略掩码+简单的文本提示，就能够丝滑实现连续的图片编辑。【图1】<br><br>这个方法主要进行了两大创新：一是支持粗略掩码输入，在保留现有内容的同时实现新元素的自然融合，其次是通过分层记忆机制实现多次修改间的编辑效果一致性。<br><br>无论是增加物体，还是删除物体，都展现出了不错的效果，能够保持前景完整与背景协调。【图2、图3】<br><br>在背后支撑着这一效果的，是一个整合了背景一致性引导和多查询解耦交叉注意力两大核心技术的框架，能够以极其简单的用户操作实现迭代式图像编辑。【图4】<br><br>- 背景一致性引导（Background Consistency Guidance）<br>利用层记忆中的信息，仅更新被编辑区域，保持未编辑区域的稳定性，同时减少计算开销。<br><br>- 多查询解耦交叉注意力（Multi-Query Disentangled Cross-Attention）<br>通过解耦查询和潜在特征的注意力，确保新对象能够自然地融入已有内容，同时保持背景和空间关系。<br><br>团队还构建了一个新的基准数据集MultiEdit-Bench，用于评估多步编辑任务中的语义对齐和空间一致性。<br><br>对这个方法感兴趣？欢迎点击下方链接查看更多内容～<br>论文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fpdf%2F2505.01079" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>项目主页：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fcarpedkm.github.io%2Fprojects%2Fimproving_edit%2Findex.html" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i15lhxc56fj31hg0ps4qp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i15li2816kj31p40tm4qp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i15li3ha7dj30t80fenbq.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i15li4vsbvj310e0hb15v.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

韩国团队开发了一种改进的AI图像编辑方法，通过粗略掩码和文本提示实现流畅的多轮图片修改。该技术采用两大创新：分层记忆机制保持编辑一致性，以及背景一致性引导技术确保未编辑区域稳定。核心框架整合了多查询解耦交叉注意力，使新元素能自然融入现有内容。团队还创建了MultiEdit-Bench数据集评估多步编辑效果。该方法支持增删物体，保持前景完整与背景协调，显著提升了AI图像编辑的稳定性和灵活性。相关论文和项目已公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T07:03:42Z
- **目录日期**: 2025-05-06
