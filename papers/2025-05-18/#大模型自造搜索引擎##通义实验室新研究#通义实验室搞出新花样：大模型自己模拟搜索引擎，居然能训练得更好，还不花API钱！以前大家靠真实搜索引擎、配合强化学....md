# #大模型自造搜索引擎##通义实验室新研究#通义实验室搞出新花样：大模型自己模拟搜索引擎，居然能训练得更好，还不花API钱！以前大家靠真实搜索引擎、配合强化学...

**URL**: https://weibo.com/6105753431/PseqE8dkd

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E8%87%AA%E9%80%A0%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E8%87%AA%E9%80%A0%E6%90%9C%E7%B4%A2%E5%BC%95%E6%93%8E%23" data-hide=""><span class="surl-text">#大模型自造搜索引擎#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%80%9A%E4%B9%89%E5%AE%9E%E9%AA%8C%E5%AE%A4%E6%96%B0%E7%A0%94%E7%A9%B6%23&amp;extparam=%23%E9%80%9A%E4%B9%89%E5%AE%9E%E9%AA%8C%E5%AE%A4%E6%96%B0%E7%A0%94%E7%A9%B6%23" data-hide=""><span class="surl-text">#通义实验室新研究#</span></a><br><br>通义实验室搞出新花样：大模型自己模拟搜索引擎，居然能训练得更好，还不花API钱！<br><br>以前大家靠真实搜索引擎、配合强化学习训练模型，出现了不少问题：<br><br>文档质量不好、成本太高、模型扩展难。<br><br>阿里团队给出的解法叫“ZeroSearch”——不对接真实搜索引擎，只靠小模型自己生成“检索结果”来训练。<br><br>一是轻量微调：给模型一点点有标签的数据，它就能生成两种文档——有用的和噪声干扰的。  <br><br>二是课程式抗噪训练：一开始喂模型高质量文档，后面逐渐掺水（加入噪声），让它逐步适应更复杂的检索任务。  <br><br>三是强化学习闭环：通过PPO和GRPO等算法，让模型自己玩转“检索-推理”流程，API成本直接归零。<br><br>实验发现，在多个问答数据集上，无论是简单的单跳问题还是复杂的多跳问题，ZeroSearch都比传统方法更厉害。特别是在用3B甚至7B的小模型时，效果就能媲美甚至超越谷歌搜索。 <a href="https://weibo.com/ttarticle/p/show?id=2309405167286524182570" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">通义实验室新研究：大模型自造搜索引擎，训练得更好，还不花API钱！</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3ly1i1ifor1v69j30oc0dpta9.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

阿里通义实验室提出"ZeroSearch"新方法，让大模型自主模拟搜索引擎进行训练。该方法通过轻量微调生成含噪声的文档数据，采用课程学习逐步增加噪声干扰，并利用强化学习形成检索-推理闭环。相比传统依赖真实搜索引擎的方法，ZeroSearch在3B/7B小模型上即实现媲美谷歌搜索的效果，且无需API调用成本。实验显示，该方法在单跳和多跳问答任务中均优于传统方案，解决了文档质量差、训练成本高和模型扩展难三大痛点。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-18T17:03:58Z
- **目录日期**: 2025-05-18
