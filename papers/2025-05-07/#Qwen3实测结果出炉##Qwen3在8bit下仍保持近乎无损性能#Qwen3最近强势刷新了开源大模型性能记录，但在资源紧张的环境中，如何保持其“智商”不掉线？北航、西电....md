# #Qwen3实测结果出炉##Qwen3在8bit下仍保持近乎无损性能#Qwen3最近强势刷新了开源大模型性能记录，但在资源紧张的环境中，如何保持其“智商”不掉线？北航、西电...

**URL**: https://weibo.com/6105753431/PqInomJtY

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Qwen3%E5%AE%9E%E6%B5%8B%E7%BB%93%E6%9E%9C%E5%87%BA%E7%82%89%23&amp;extparam=%23Qwen3%E5%AE%9E%E6%B5%8B%E7%BB%93%E6%9E%9C%E5%87%BA%E7%82%89%23" data-hide=""><span class="surl-text">#Qwen3实测结果出炉#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Qwen3%E5%9C%A88bit%E4%B8%8B%E4%BB%8D%E4%BF%9D%E6%8C%81%E8%BF%91%E4%B9%8E%E6%97%A0%E6%8D%9F%E6%80%A7%E8%83%BD%23&amp;extparam=%23Qwen3%E5%9C%A88bit%E4%B8%8B%E4%BB%8D%E4%BF%9D%E6%8C%81%E8%BF%91%E4%B9%8E%E6%97%A0%E6%8D%9F%E6%80%A7%E8%83%BD%23" data-hide=""><span class="surl-text">#Qwen3在8bit下仍保持近乎无损性能#</span></a><br><br>Qwen3最近强势刷新了开源大模型性能记录，但在资源紧张的环境中，如何保持其“智商”不掉线？<br><br>北航、西电与ETH的联合团队首次系统性评估了Qwen3在后训练量化（PTQ）下的鲁棒性。<br><br>他们用5种经典PTQ方法，测试了1到8比特位宽在多个任务下的表现，还拉上了LLaMA3来做对比。<br><br>结果显示，在8比特设置下，Qwen3性能几乎无损，非常适合实际部署。<br><br>但随着位宽降低到4比特以下，性能开始明显下滑，到了2比特更是“惨不忍睹”。尽管某些方法如GPTQ、AWQ能稍微挽救一下，但整体仍需新技术突破。<br><br>此外，研究还发现Qwen3对激活量化格外敏感，哪怕是常用的SmoothQuant方法，也难免带来性能衰减。<br><br>而且参数越大的模型，在量化下越稳定，这可能是因为大模型有更强的“容错力”。<br><br>不过跟LLaMA3一比，Qwen3在低位宽下表现差距就很明显了，特别是在超低比特场景中，这可能是因为Qwen3预训练过程更彻底，冗余更少，对量化引发的信息损失也更敏感。<br><br>详细评估过程：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FL9-w3ngZvo1I7zIqbHhD6w" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">首次系统评估Qwen3在后训练量化下的鲁棒性：8bit下仍保持近乎无损性能</span></a><br>[图片]<br>[图片]<br>OpenAI 达成协议，以 30 亿美元收购初创公司 Windsurf<br><a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.bloomberg.com%2Fnews%2Farticles%2F2025-05-06%2Fopenai-reaches-agreement-to-buy-startup-windsurf-for-3-billion%3Fembedded-checkout%3Dtrue" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i16wjtqtq8j30n40fuaeu.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i16wjuwkq0j30n40ggjwt.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

北航、西电与ETH联合团队对Qwen3大模型的后训练量化(PTQ)性能进行了系统评估。测试显示，Qwen3在8bit量化下性能几乎无损，适合实际部署；但4bit以下性能明显下降，2bit时表现较差。相比LLaMA3，Qwen3在低位宽下表现更敏感，可能因其预训练更彻底、冗余更少。研究还发现大模型对量化更稳定，且Qwen3对激活量化特别敏感。该评估为开源大模型在资源受限环境中的部署提供了重要参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T09:04:21Z
- **目录日期**: 2025-05-07
