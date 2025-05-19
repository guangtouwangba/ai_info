# #Transformer的传统方法解释##用N-gram预测LLM行为#Transformer的行为，竟能被十几年前的老方法——N-gram猜个八九不离十？DeepMind这篇论文，署名只有一个作者...

**URL**: https://weibo.com/6105753431/Psy494FiT

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Transformer%E7%9A%84%E4%BC%A0%E7%BB%9F%E6%96%B9%E6%B3%95%E8%A7%A3%E9%87%8A%23&amp;extparam=%23Transformer%E7%9A%84%E4%BC%A0%E7%BB%9F%E6%96%B9%E6%B3%95%E8%A7%A3%E9%87%8A%23" data-hide=""><span class="surl-text">#Transformer的传统方法解释#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8N-gram%E9%A2%84%E6%B5%8BLLM%E8%A1%8C%E4%B8%BA%23&amp;extparam=%23%E7%94%A8N-gram%E9%A2%84%E6%B5%8BLLM%E8%A1%8C%E4%B8%BA%23" data-hide=""><span class="surl-text">#用N-gram预测LLM行为#</span></a><br><br>Transformer的行为，竟能被十几年前的老方法——N-gram猜个八九不离十？<br><br>DeepMind这篇论文，署名只有一个作者（Timothy Nguyen），而且还是个哲学教授。<br><br>该论文用非常“老派”的方法重新审视LLM的预测机制，结果发现，LLM在很多时候的预测结果，其实可以被简单的N-gram规则逼近。<br><br>具体来说，作者设计了三种 N-gram 规则：<br><br>- suffix：只看最后几个 token；<br>- subgram：选上下文里的几个关键 token；<br>- all：上下文全都用上。<br><br>这三类规则越往后越复杂，用来“猜”下一个 token，和模型预测做比对。结果发现：<br><br>1、在 TinyStories（一个小体量但结构清晰的文本数据集）上，有79%的预测，是能被这些规则解释的；  <br><br>2、换成复杂很多的 Wikipedia，准确率也还有 68%。<br><br>你可以理解为，大模型有相当一部分行为，其实只是“在看见过的上下文里重复出现”。<br><br>为了防止这是巧合，论文还干了几件事：<br><br>- 不同模型规模（160M 和 1.4B 参数）都试了；<br>- 不同数据集（TinyStories 和 Wikipedia）都试了；<br>- 分析了上下文的频率、模型预测的波动程度、规则覆盖率之间的关系。<br><br>结果显示，如果一个上下文本身很频繁、模型在这个上下文下预测很稳定（也就是 低方差），那么这些N-gram规则就特别好用，能高度还原模型的行为。<br><br>作者还提出一个有意思的发现：只用训练数据，就能监测模型是否开始过拟合。具体操作是：<br><br>- 把上下文简化一下（比如只保留几个 token）；<br>- 看模型对这些“缩水输入”的表现有没有明显下降。<br><br>如果下降得特别快，说明模型已经开始过度依赖上下文细节了，也就是在过拟合。<br><br>文章表示，Transformer在训练初期，主要靠简单的 N-gram 规则来做预测，随着训练推进，才逐渐掌握更复杂的结构模式。也就是说，大模型在训练的时候，自己摸索出了学习节奏，一开始学简单的，熟练后再学难的。<br><br>这种从“简单规则”过渡到“复杂规则”的过程，非常像人类的课程式学习（curriculum learning）。<br><br>当然，这不意味着我们可以靠N-gram去重建一个 LLM，只是说，在某些上下文情况下，模型的行为很像“查表”，但查的是训练数据里的“重复模板”。<br><br>网友看到这篇文章之后感叹：“我以前就是这么理解语言模型的，后来好像主流关键都不流行这么看了。”<br><br>感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2407.12034" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1kugpq7whj30j70p0q99.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1kugqus1rj30if0p00yh.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

DeepMind研究发现，Transformer模型的行为可用传统N-gram方法解释。论文通过三种N-gram规则（suffix/subgram/all）测试发现：在TinyStories数据集上79%的预测、Wikipedia上68%的预测能被这些简单规则逼近。研究表明，模型初期依赖N-gram规则，后期才学习复杂模式，类似人类课程学习。当上下文高频且预测稳定时，N-gram效果最佳。该方法还能通过简化输入检测过拟合。虽然不能替代LLM，但揭示了模型部分行为本质是"重复训练数据中的模式"。论文验证了不同规模模型和数据集均存在此现象。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T09:03:26Z
- **目录日期**: 2025-05-19
