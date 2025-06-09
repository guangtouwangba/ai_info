# #图神经网络迎来新突破##思维链推理刷新图学习表现上限#思维链竟能用在图数据上？中科大等团队提出了一个叫GCoT的新框架，它把语言模型里的“思维链”方法搬到了...

**URL**: https://weibo.com/6105753431/PvAcR5kiw

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%9B%BE%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C%E8%BF%8E%E6%9D%A5%E6%96%B0%E7%AA%81%E7%A0%B4%23&amp;extparam=%23%E5%9B%BE%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C%E8%BF%8E%E6%9D%A5%E6%96%B0%E7%AA%81%E7%A0%B4%23" data-hide=""><span class="surl-text">#图神经网络迎来新突破#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%80%9D%E7%BB%B4%E9%93%BE%E6%8E%A8%E7%90%86%E5%88%B7%E6%96%B0%E5%9B%BE%E5%AD%A6%E4%B9%A0%E8%A1%A8%E7%8E%B0%E4%B8%8A%E9%99%90%23&amp;extparam=%23%E6%80%9D%E7%BB%B4%E9%93%BE%E6%8E%A8%E7%90%86%E5%88%B7%E6%96%B0%E5%9B%BE%E5%AD%A6%E4%B9%A0%E8%A1%A8%E7%8E%B0%E4%B8%8A%E9%99%90%23" data-hide=""><span class="surl-text">#思维链推理刷新图学习表现上限#</span></a><br><br>思维链竟能用在图数据上？<br><br>中科大等团队提出了一个叫GCoT的新框架，它把语言模型里的“思维链”方法搬到了图数据上。这是首次尝试在没有文本信息的图数据上搞链式推理。<br><br>图数据结构复杂、缺少文字，传统提示学习方法在这儿很难发力。GCoT的核心点就是“拆步骤”：<br><br>- 首先输入图和提示，预训练模型做初步推断；<br>    <br>- 然后融合多层隐藏表示，生成一份叫“思维”的中间表达；<br>    <br>- 接着用这个思维来生成节点专属提示，继续下一步推理。<br>    <br>为了验证效果，研究团队在八个数据集上测了 GCoT 的表现。结果在1到5个样本这么少的情况下，GCoT还是稳压其他SOTA模型。 <a href="https://weibo.com/ttarticle/p/show?id=2309405175273494020112" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">8个数据集全面胜出！思维链推理刷新图学习表现上限</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3ly1i27x25yo00j30gw09iwfi.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

中科大团队提出GCoT框架，首次将语言模型的思维链推理应用于无文本的图数据。该方法通过三步走：初步推断→生成中间"思维"表示→创建节点专属提示进行迭代推理。在8个数据集测试中，GCoT仅用1-5个样本就超越现有最优模型，突破了图神经网络在少样本场景下的性能上限。该创新解决了图数据结构复杂、缺乏文本信息导致的传统提示学习困难问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T01:31:09Z
- **目录日期**: 2025-06-09
