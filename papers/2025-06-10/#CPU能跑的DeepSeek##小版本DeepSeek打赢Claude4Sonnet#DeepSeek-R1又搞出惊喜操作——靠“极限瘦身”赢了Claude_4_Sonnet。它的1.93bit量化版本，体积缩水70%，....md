# #CPU能跑的DeepSeek##小版本DeepSeek打赢Claude4Sonnet#DeepSeek-R1又搞出惊喜操作——靠“极限瘦身”赢了Claude 4 Sonnet。它的1.93bit量化版本，体积缩水70%，...

**URL**: https://weibo.com/6105753431/PvTO78PTM

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23CPU%E8%83%BD%E8%B7%91%E7%9A%84DeepSeek%23&amp;extparam=%23CPU%E8%83%BD%E8%B7%91%E7%9A%84DeepSeek%23" data-hide=""><span class="surl-text">#CPU能跑的DeepSeek#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%B0%8F%E7%89%88%E6%9C%ACDeepSeek%E6%89%93%E8%B5%A2Claude4Sonnet%23&amp;extparam=%23%E5%B0%8F%E7%89%88%E6%9C%ACDeepSeek%E6%89%93%E8%B5%A2Claude4Sonnet%23" data-hide=""><span class="surl-text">#小版本DeepSeek打赢Claude4Sonnet#</span></a><br><br>DeepSeek-R1又搞出惊喜操作——靠“极限瘦身”赢了Claude 4 Sonnet。<br><br>它的1.93bit量化版本，体积缩水70%，但在aider编程榜上却打败了Sonnet和自己1月的满血版本，成了“轻装上阵也能赢”的代表。<br><br>来看看具体怎么回事：<br><br>- 榜单成绩亮眼：R1-0528（量化版）在aider榜拿到60%，击败Claude 4 Sonnet的56.4%，比1月的R1还强。而满血R1-0528更是直接打到71.4%，超过不开“思考”的Claude 4 Opus。<br>    <br>- 量化不用GPU也能跑：这个奇迹来自Unsloth团队，他们做了从1.66bit到5.5bit共9个版本。比如1.78bit版本，只要64GB内存，就能在CPU上跑，每秒输出1个token；24G显卡+128G内存能跑到5个token/s。<br>    <br>- 团队战绩强悍：Unsloth不仅做了DeepSeek，连Qwen、Phi、Mistral、Llama也都被他们优化过，最快提速达50%、最省内存砍半。GitHub上已有4万星星。<br>    <br>- 连打游戏都赢了：Hao AI Lab用R1-0528测试了一波人类小游戏，包括：<br>    - 俄罗斯方块：坚持最久，打败o4-mini，仅次于o3；<br>    - 2048、推箱子：表现大幅领先1月版；<br>    - 糖果传奇：得分548，领先o4-mini近20分，仅次于o3。<br>        <br>最后再来个推荐版本和跑法：<br>    - 最推荐：2.4bit和2.7bit，精度&amp;体积更均衡；<br>    - 跑得动的关键：下载文件大小 ≤ 显存+内存总量；<br>    - 想稳妥运行，Unsloth建议180GB统一内存或RAM+显存合计超180GB。<br><a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F3AcTgLHBIANjLGeV5HyMmQ" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">1.93bit版DeepSeek-R1编程超过Claude 4 Sonnet，不用GPU也能运行</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2a8hdiqwfj30zk0cc415.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2a8h302ftg30lu0lo1kz.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2a8gpbz2fj30zk0n4k2u.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2a8gvlfkcj30zk0l7nem.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

DeepSeek-R1通过1.93bit极限量化实现70%体积压缩，在aider编程榜以60%得分击败Claude 4 Sonnet（56.4%），其完整版更达71.4%。Unsloth团队开发的量化版本支持CPU运行（如1.78bit版仅需64GB内存，1 token/s），同时优化多款主流模型。实测显示，该模型在俄罗斯方块、2048等游戏中表现优异，推荐2.4-2.7bit平衡版，运行需180GB内存或显存合计。GitHub获4万星，突显其高效轻量化优势。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T09:03:52Z
- **目录日期**: 2025-06-10
