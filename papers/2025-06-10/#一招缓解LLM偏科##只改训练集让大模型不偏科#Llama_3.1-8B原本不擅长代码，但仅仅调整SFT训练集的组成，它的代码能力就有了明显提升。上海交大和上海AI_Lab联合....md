# #一招缓解LLM偏科##只改训练集让大模型不偏科#Llama 3.1-8B原本不擅长代码，但仅仅调整SFT训练集的组成，它的代码能力就有了明显提升。上海交大和上海AI Lab联合...

**URL**: https://weibo.com/6105753431/PvTXgf4zt

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E6%8B%9B%E7%BC%93%E8%A7%A3LLM%E5%81%8F%E7%A7%91%23&amp;extparam=%23%E4%B8%80%E6%8B%9B%E7%BC%93%E8%A7%A3LLM%E5%81%8F%E7%A7%91%23" data-hide=""><span class="surl-text">#一招缓解LLM偏科#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8F%AA%E6%94%B9%E8%AE%AD%E7%BB%83%E9%9B%86%E8%AE%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%E4%B8%8D%E5%81%8F%E7%A7%91%23&amp;extparam=%23%E5%8F%AA%E6%94%B9%E8%AE%AD%E7%BB%83%E9%9B%86%E8%AE%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%E4%B8%8D%E5%81%8F%E7%A7%91%23" data-hide=""><span class="surl-text">#只改训练集让大模型不偏科#</span></a><br><br>Llama 3.1-8B原本不擅长代码，但仅仅调整SFT训练集的组成，它的代码能力就有了明显提升。<br><br>上海交大和上海AI Lab联合提出了IDEAL方法，核心思路很直接：别一味增加数据量，而是科学地调整不同领域数据的比例。<br><br>研究团队发现：<br><br>- SFT阶段的数据量不是越多越好。<br>    <br>- 数据配比不合适时，多反而可能“反伤”模型。<br>    <br>- 有些能力甚至会在SFT后退化。<br><br>IDEAL方法背后的机制是，通过理论建模，评估各领域数据对模型验证集性能的影响；再用K-FAC技术高效估算Hessian逆矩阵，减轻计算压力。最后，通过控制超参数m调整配比，以获得更平衡的训练集。 <a href="https://weibo.com/ttarticle/p/show?id=2309405176032621101495" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">一招缓解LLM偏科！调整训练集组成，“秘方”在此</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i2abtk7694j30kg0bi3zl.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

上海交大与上海AI Lab联合提出IDEAL方法，通过科学调整SFT训练集的数据配比（而非单纯增加数据量）显著提升Llama 3.1-8B的代码能力。研究发现：1）数据量过多可能损害模型性能；2）不当配比会导致某些能力退化。该方法通过理论建模评估领域数据影响，利用K-FAC技术高效计算Hessian逆矩阵，并引入超参数m动态优化数据比例，实现更均衡的模型表现。实验显示仅调整训练集组成即可有效缓解大模型"偏科"问题。（98字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T23:03:53Z
- **目录日期**: 2025-06-10
