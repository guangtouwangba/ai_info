# #一图区分Transformer和MoE##MoE是分工协作版的Transformer#还不了解Transformer和MoE两种架构的不同之处？别怕，这张动图让一切一目了然⬇️Transformer和MoE的...

**URL**: https://weibo.com/6105753431/PoraBuMh4

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E5%9B%BE%E5%8C%BA%E5%88%86Transformer%E5%92%8CMoE%23&amp;extparam=%23%E4%B8%80%E5%9B%BE%E5%8C%BA%E5%88%86Transformer%E5%92%8CMoE%23" data-hide=""><span class="surl-text">#一图区分Transformer和MoE#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23MoE%E6%98%AF%E5%88%86%E5%B7%A5%E5%8D%8F%E4%BD%9C%E7%89%88%E7%9A%84Transformer%23&amp;extparam=%23MoE%E6%98%AF%E5%88%86%E5%B7%A5%E5%8D%8F%E4%BD%9C%E7%89%88%E7%9A%84Transformer%23" data-hide=""><span class="surl-text">#MoE是分工协作版的Transformer#</span></a><br><br>还不了解Transformer和MoE两种架构的不同之处？别怕，这张动图让一切一目了然⬇️<br><br>Transformer和MoE的主要区别出现在解码层当中，经过层归一化后，MoE取代传统Transformer的前馈神经网络（Feedforward Neural Network，FNN），在模型中引入了稀疏性。<br><br>具体来说：<br><br>Transformer：使用传统的前馈神经网络，所有参数都会被激活，如同全班同学一起做同一道数学题。<br><br>MoE通常由两个核心部分构成：<br>1. 路由器或网关网络（Router）：路由器将会决定如何分配Token，选择特定的专家。<br>2. 专家模块：由多个专家组成，每个专家往往本身即是完整的前馈网络。<br><br>一句话总结：MoE是Transformer的“分工协作版”，适合更复杂的任务。<img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0pl5o23ntg30v00vadt3.gif" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Transformer和MoE的主要区别在于解码层结构：Transformer使用全激活的前馈神经网络(FNN)，而MoE采用稀疏架构。MoE包含路由器(Router)和专家模块两部分，路由器分配Token给特定专家(每个专家本身是独立FNN)，实现参数选择性激活。相比Transformer"全班一起做题"的工作模式，MoE通过专家分工协作，在保持模型规模的同时提升计算效率，适合处理更复杂任务。这种设计使MoE成为Transformer的高效变体，核心创新在于用动态路由机制替代了传统的全连接前馈网络。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T09:04:12Z
- **目录日期**: 2025-04-22
