# #不用RL也能搞定大模型微调##只需一条样本做无监督训练#强化学习太烧钱？试试只用一条样本、训练十步就见效的新方法。来自Ubiquant的研究团队提出了“熵最小化”...

**URL**: https://weibo.com/6105753431/Puws3fkK6

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%8D%E7%94%A8RL%E4%B9%9F%E8%83%BD%E6%90%9E%E5%AE%9A%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%BE%AE%E8%B0%83%23&amp;extparam=%23%E4%B8%8D%E7%94%A8RL%E4%B9%9F%E8%83%BD%E6%90%9E%E5%AE%9A%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%BE%AE%E8%B0%83%23" data-hide=""><span class="surl-text">#不用RL也能搞定大模型微调#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8F%AA%E9%9C%80%E4%B8%80%E6%9D%A1%E6%A0%B7%E6%9C%AC%E5%81%9A%E6%97%A0%E7%9B%91%E7%9D%A3%E8%AE%AD%E7%BB%83%23&amp;extparam=%23%E5%8F%AA%E9%9C%80%E4%B8%80%E6%9D%A1%E6%A0%B7%E6%9C%AC%E5%81%9A%E6%97%A0%E7%9B%91%E7%9D%A3%E8%AE%AD%E7%BB%83%23" data-hide=""><span class="surl-text">#只需一条样本做无监督训练#</span></a><br><br>强化学习太烧钱？试试只用一条样本、训练十步就见效的新方法。<br><br>来自Ubiquant的研究团队提出了“熵最小化”（Entropy Minimization，EM）方法，能让大语言模型快速“变聪明”——<br><br>不需要标注数据、不用奖励模型，只靠模型自己对输出的置信度进行调整，就能在数学推理任务上显著提升性能。<br><br>相比传统RL，EM的优势在于：<br><br>- 完全无监督训练，不需要奖励函数；<br><br>- 数据极简，只用一条高质量示例；<br><br>- 训练极快，10步内就能提升准确率；<br><br>- 表现可媲美甚至超越使用成千上万条数据的RL方法。<br><br>EM的做法简单粗暴：让模型生成答案时，把预测分布的“熵”压低，也就是逼它只挑最自信的答案。比如一个问题它原本有点犹豫，现在被训练得“非它莫属”，从而提高准确率。<br><br>研究还发现，选择哪一条样本非常关键。团队通过观察模型对问题的回答是否稳定来挑样本——不稳定说明“纠结”，最适合训练EM。实际只用一道物理题，就让Qwen2.5-Math-7B模型的数学推理能力暴涨。 <a href="https://weibo.com/ttarticle/p/show?id=2309405172745871950031" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">挑战强化学习后训练霸权！全新无监督方法仅需1条数据+10步优化</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i1zuqc7orvj30om0duq4o.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Ubiquant研究团队提出"熵最小化"(EM)方法，无需强化学习(RL)即可高效微调大模型。该方法完全无监督，仅需1条高质量样本和10步训练即可显著提升模型性能。核心原理是通过降低输出分布的熵值，迫使模型选择最自信的答案。实验显示，在数学推理任务上，EM方法效果媲美甚至超越需要大量数据的RL方法。关键创新在于选择模型回答不稳定的问题作为训练样本，例如仅用1道物理题就使Qwen2.5-Math-7B模型性能大幅提升。该方法具有低成本、高效率的优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-01T20:02:46Z
- **目录日期**: 2025-06-01
