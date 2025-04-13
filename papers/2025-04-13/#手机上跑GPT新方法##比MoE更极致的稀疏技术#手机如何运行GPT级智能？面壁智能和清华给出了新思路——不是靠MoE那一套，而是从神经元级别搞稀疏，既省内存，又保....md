# #手机上跑GPT新方法##比MoE更极致的稀疏技术#手机如何运行GPT级智能？面壁智能和清华给出了新思路——不是靠MoE那一套，而是从神经元级别搞稀疏，既省内存，又保...

**URL**: https://weibo.com/6105753431/PmWw61M8f

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%89%8B%E6%9C%BA%E4%B8%8A%E8%B7%91GPT%E6%96%B0%E6%96%B9%E6%B3%95%23&amp;extparam=%23%E6%89%8B%E6%9C%BA%E4%B8%8A%E8%B7%91GPT%E6%96%B0%E6%96%B9%E6%B3%95%23" data-hide=""><span class="surl-text">#手机上跑GPT新方法#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%AF%94MoE%E6%9B%B4%E6%9E%81%E8%87%B4%E7%9A%84%E7%A8%80%E7%96%8F%E6%8A%80%E6%9C%AF%23&amp;extparam=%23%E6%AF%94MoE%E6%9B%B4%E6%9E%81%E8%87%B4%E7%9A%84%E7%A8%80%E7%96%8F%E6%8A%80%E6%9C%AF%23" data-hide=""><span class="surl-text">#比MoE更极致的稀疏技术#</span></a><br><br>手机如何运行GPT级智能？面壁智能和清华给出了新思路——<br><br>不是靠MoE那一套，而是从神经元级别搞稀疏，既省内存，又保证性能。<br><br>这项技术名为CFM（Configurable Foundation Models），核心思路是用模型自身天然的稀疏激活特性，不走MoE那种“top k专家”路线，而是让神经元自由发挥：任务难就多激活几个，任务简单就偷懒省电。<br><br>相比之下，MoE训练时得限制激活专家数量，不然一不小心就算不动了。CFM不需要这些“平衡负载”的麻烦，全量参数都能参与运算，省心又高效。<br><br>CFM另一大亮点是灵活组合模型模块，就像拼积木那样，任务导向建模更高效。作者肖朝军提到，这种神经元级稀疏在FFN层尤其有效，特别适合端侧场景，比如手机、笔记本、智能家居等。<br> <a href="https://weibo.com/ttarticle/p/show?id=2309405154681583173676" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">手机实现GPT级智能，比MoE更极致的稀疏技术：省内存效果不减｜对话面</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0e83zhm8lj30r70fb76u.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

面壁智能与清华团队提出CFM技术，通过神经元级稀疏激活实现在手机等端侧设备运行GPT级模型。相比MoE固定激活专家数量的方式，CFM根据任务难度动态调整激活神经元数量，既节省内存又保持性能。该技术特别适用于FFN层，支持灵活组合模型模块，无需负载平衡处理。CFM充分发挥模型天然稀疏特性，使全量参数参与运算，为手机、笔记本等设备提供高效AI解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T21:03:23Z
- **目录日期**: 2025-04-13
