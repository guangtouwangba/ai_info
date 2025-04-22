# #一图区分Transformer和MoE##MoE是分工协作版的Transformer#还不了解Transformer和MoE两种架构的不同之处？别怕，这张动图让一切一目了然⬇️Transformer和MoE的...

**URL**: https://weibo.com/6105753431/PoraBuMh4

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E5%9B%BE%E5%8C%BA%E5%88%86Transformer%E5%92%8CMoE%23&amp;extparam=%23%E4%B8%80%E5%9B%BE%E5%8C%BA%E5%88%86Transformer%E5%92%8CMoE%23" data-hide=""><span class="surl-text">#一图区分Transformer和MoE#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23MoE%E6%98%AF%E5%88%86%E5%B7%A5%E5%8D%8F%E4%BD%9C%E7%89%88%E7%9A%84Transformer%23&amp;extparam=%23MoE%E6%98%AF%E5%88%86%E5%B7%A5%E5%8D%8F%E4%BD%9C%E7%89%88%E7%9A%84Transformer%23" data-hide=""><span class="surl-text">#MoE是分工协作版的Transformer#</span></a><br><br>还不了解Transformer和MoE两种架构的不同之处？别怕，这张动图让一切一目了然⬇️<br><br>Transformer和MoE的主要区别出现在解码层当中，经过层归一化后，MoE取代传统Transformer的前馈神经网络（Feedforward Neural Network，FNN），在模型中引入了稀疏性。<br><br>具体来说：<br><br>Transformer：使用传统的前馈神经网络，所有参数都会被激活，如同全班同学一起做同一道数学题。<br><br>MoE通常由两个核心部分构成：<br>1. 路由器或网关网络（Router）：路由器将会决定如何分配Token，选择特定的专家。<br>2. 专家模块：由多个专家组成，每个专家往往本身即是完整的前馈网络。<br><br>一句话总结：MoE是Transformer的“分工协作版”，适合更复杂的任务。<img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0pl5o23ntg30v00vadt3.gif" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Transformer和MoE的主要区别在于解码层结构：Transformer使用全激活的前馈神经网络(FNN)，所有参数参与计算；而MoE采用稀疏激活机制，通过路由器(Router)将输入分配给特定的专家模块(前馈网络)。MoE就像分工协作的Transformer版本，每个输入只激活部分专家，提高了模型效率，适合处理更复杂的任务。这种架构既保持了Transformer的核心能力，又通过专家分工实现了计算资源的优化利用。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T08:04:11Z
- **目录日期**: 2025-04-22
